// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"io/ioutil"
	"net/http"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

// JSONparser defines a parser that converts string bytes into an object
type JSONparser func([]byte) interface{}

// JSONhandler is a handler that can take a JSON input data and returns JSON output data
//
//   inputData -- a parsed JSON structure [may be nil]
//
type JSONhandler func(w http.ResponseWriter, r *http.Request, inputData interface{}) (results *Results)

// Jhandler makes a handler wrapping another handler in order to send JSON response
// Optionally, this handler may parse input JSON
//
//   parser -- a parse for input JSON data [may be nil]
//
//   NOTE: (1) works only with POST and GET methods
//         (2) allow CORS
//         (3) Jhandler catch panics already => no need to wrap with Ehandler
//             Will return a JSON object with "success":false and "error":message on errors/panicks
//         (4) if the hfcn handler does not return a "Results" object, the returned JSON
//             will be successful but not authorized
//
func Jhandler(hfcn JSONhandler, parser JSONparser) http.HandlerFunc {

	// return standard handler
	return func(w http.ResponseWriter, r *http.Request) {

		// catch errors
		defer func() {
			if err := recover(); err != nil {
				lio.Ff(w, RjsonFailed(err))
			}
		}()

		// skip other methods
		if r.Method != "POST" && r.Method != "GET" {
			check.Panic("Jhandler works with POST and GET only\n")
		}

		// set access and return type
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Content-Type", "application/json")

		// auxiliary
		var inputData interface{}

		// with input JSON
		if parser != nil {

			// auxiliary
			var inputJSON []byte

			// handle POST
			var err error
			if r.Method == "POST" {
				inputJSON, err = ioutil.ReadAll(r.Body) // parse 'body'
				if err != nil {
					check.Panic("ReadAll failed: %v\n", err)
				}
			}

			// handle GET
			if r.Method == "GET" {
				err = r.ParseForm()
				if err != nil {
					check.Panic("cannot parse form: %v\n", err)
				}
				for str := range r.Form {
					inputJSON = []byte(str)
					break
				}
			}

			// parse json
			inputData = parser(inputJSON)
		}

		// call handler
		results := hfcn(w, r, inputData)
		if results == nil {
			lio.Ff(w, RjsonUnauthorized())
			return
		}

		// set results
		results.Success = true
		lio.Ff(w, results.ToJSON())
	}
}
