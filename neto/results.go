// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"encoding/json"
	"net/http"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

// Results holds data to be constructed as a JSON response
type Results struct {
	Authorized bool                   // indicates whether user is authorized or not; will become "authorized" member of JSON data
	Success    bool                   // indicates if request performed successfully
	Data       map[string]interface{} // contains aditional data
}

// NewResults returns new object
func NewResults() (o *Results) {
	o = new(Results)
	return
}

// NewResultsFromJSON returns new object from JSON data
func NewResultsFromJSON(data []byte) (o *Results) {
	o = new(Results)
	err := json.Unmarshal(data, o)
	if err != nil {
		check.Panic("cannot parse JSON data: %v\n", err)
	}
	return
}

// Set sets data
func (o *Results) Set(key string, value interface{}) {
	if len(o.Data) == 0 {
		o.Data = make(map[string]interface{})
	}
	o.Data[key] = value
}

// ToJSON converts Results to ToJSON string
//   Example: returns `{"authorized":true,"success":true}`
func (o *Results) ToJSON() string {

	// empty map
	if len(o.Data) == 0 {
		return lio.Sf(`{"authorized":%v,"success":%v}`, o.Authorized, o.Success)
	}

	// has map
	o.Data["authorized"] = o.Authorized
	o.Data["success"] = o.Success
	b, err := json.Marshal(o.Data)
	if err != nil {
		return RjsonFailed(err)
	}
	return string(b)
}

// RjsonFailed makes a JSON string representing a failed "Results", with error
//   Example: returns `{"authorized":false,"success":false,"error":"error message"}`
//   Remember to: w.Header().Set("Content-Type", "application/json")
func RjsonFailed(err interface{}) string {
	return lio.Sf(`{"authorized":false,"success":false,"error":%q}`, err)
}

// RjsonUnauthorized makes a JSON string representing an unauthorized "Results",
// but during a successful request
//   Example: returns `{"authorized":false,"success":true}`
//   Remember to: w.Header().Set("Content-Type", "application/json")
func RjsonUnauthorized() string {
	return `{"authorized":false,"success":true}`
}

// RjsonAllGood makes an JSON string representing an authorized and successful "Results"
//   Example: returns `{"authorized":true,"success":true}`
//   Remember to: w.Header().Set("Content-Type", "application/json")
func RjsonAllGood() string {
	return `{"authorized":true,"success":true}`
}

// RjsonAllGoodWithData makes an JSON string representing an authorized and successful "Results"
// with extra Data
//
//   Example: returns `{"authorized":true,"success":true,"data":{"memberA":"A","memberB":123}}`
//   Remember to: w.Header().Set("Content-Type", "application/json")
//
//   data -- should be pairs of values; e.g. "memberA", "A", "memberB", 123
//
func RjsonAllGoodWithData(data ...interface{}) (l string) {
	nargs := len(data)
	if nargs%2 != 0 {
		check.Panic("number of items in data must be even.\ndata = %+v\n", data)
	}
	l = `{"authorized":true,"success":true,"data":{`
	for i := 0; i < nargs; i += 2 {
		if i > 0 {
			l += ","
		}
		key := data[i]
		val := data[i+1]
		switch val.(type) {
		case string:
			l += lio.Sf("%q:%q", key, val)
		default:
			l += lio.Sf("%q:%v", key, val)
		}
	}
	return l + "}}"
}

// RjsonAllGoodWithDataString makes an JSON string representing an authorized and successful "Results"
// with extra Data given as a string
//
//   Example: returns `{"authorized":true,"success":true,"data":{"memberA":"A","memberB":123}}`
//   Remember to: w.Header().Set("Content-Type", "application/json")
//
//   data -- example `{"memberA":"A","memberB":123}`
//
func RjsonAllGoodWithDataString(dataString string) (l string) {
	return lio.Sf(`{"authorized":true,"success":true,"data":%s}`, dataString)
}

// ------------- writers ------------------

// WjsonFailed writes RjsonFailed
func WjsonFailed(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	lio.Ff(w, RjsonFailed(err))
}

// WjsonUnauthorized writes RjsonUnauthorized
func WjsonUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	lio.Ff(w, RjsonUnauthorized())
}

// WjsonAllGood writes RjsonAllGood
func WjsonAllGood(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	lio.Ff(w, RjsonAllGood())
}

// WjsonAllGoodWithData writes RjsonAllGoodWithData
func WjsonAllGoodWithData(w http.ResponseWriter, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	lio.Ff(w, RjsonAllGoodWithData(data...))
}

// WjsonAllGoodWithDataString writes RjsonAllGoodWithDataString
func WjsonAllGoodWithDataString(w http.ResponseWriter, dataString string) {
	w.Header().Set("Content-Type", "application/json")
	lio.Ff(w, RjsonAllGoodWithDataString(dataString))
}
