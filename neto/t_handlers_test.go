// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

// MyData is an example of data strucutre
type MyData struct {
	ID   int
	Name string
}

// jparser parses JSON string to MyData
func jparser(in []byte) interface{} {
	data := new(MyData)
	err := json.Unmarshal(in, data)
	if err != nil {
		check.Panic("cannot parse JSON data: %v\n", err)
	}
	return data
}

// jhandler01 implements a handler that taks a JSON input data and returns a JSON response
func jhandler01(w http.ResponseWriter, r *http.Request, inputData interface{}) (results *Results) {
	data := inputData.(*MyData)
	results = NewResults()
	results.Set("gotID", data.ID)
	results.Set("gotName", data.Name)
	return
}

// jhandler02 implements a handler that returns a JSON response, but do not take JSON input
func jhandler02(w http.ResponseWriter, r *http.Request, inputData interface{}) (results *Results) {
	results = NewResults()
	results.Set("hello", 123)
	return
}

// jhandler03 implements a handler that receives nothing and returns nothing
func jhandler03(w http.ResponseWriter, r *http.Request, inputData interface{}) (results *Results) {
	return
}

// jhandler04 implements a handler that panics
func jhandler04(w http.ResponseWriter, r *http.Request, inputData interface{}) (results *Results) {
	check.Panic("jhandler04 wants to stop")
	return
}

// testRequestWithJSON runs test that makes a GET or POST request with inputJSON data
func testRequestWithJSON(tst *testing.T, server *httptest.Server, method, inputJSON, correctOutput string) {

	// make request
	var response *http.Response
	var err error
	if method == "GET" {
		response, err = http.Get(server.URL + "?" + url.QueryEscape(inputJSON))
	} else {
		response, err = http.Post(server.URL, "application/json", strings.NewReader(inputJSON))
	}
	defer func() {
		if response != nil {
			response.Body.Close()
		}
	}()

	// check if request worked
	if err != nil {
		tst.Errorf("%s failed: %v\n", method, err)
		return
	}

	// check response
	if response.StatusCode != 200 {
		tst.Errorf("%s failed with Status = %v\n", method, response.Status)
		return
	}

	// check outputData
	outputData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		tst.Errorf("ioutil.ReadAll failed: %v\n", err)
		return
	}
	lio.Pf("got: %v\n", string(outputData))
	check.String(tst, lio.Sf("%s: outputData", method), string(outputData), correctOutput)
}

func TestJhandler01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Jhandler01. GET and POST with input JSON")

	// create test server
	server := httptest.NewServer(Jhandler(jhandler01, jparser))
	defer server.Close()

	// correct output
	inputJSON := `{"id":123, "name":"dorival"}`
	correctOutput := `{"authorized":false,"gotID":123,"gotName":"dorival","success":true}`

	// test
	testRequestWithJSON(tst, server, "GET", inputJSON, correctOutput)
	testRequestWithJSON(tst, server, "POST", inputJSON, correctOutput)
}

func TestJhandler02(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Jhandler02. GET and POST without input but with output")

	// create test server
	server := httptest.NewServer(Jhandler(jhandler02, nil))
	defer server.Close()

	// test
	testRequestWithJSON(tst, server, "GET", "", `{"authorized":false,"hello":123,"success":true}`)
	testRequestWithJSON(tst, server, "POST", "", `{"authorized":false,"hello":123,"success":true}`)
}

func TestJhandler03(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Jhandler03. GET and POST without input or output")

	// create test server
	server := httptest.NewServer(Jhandler(jhandler03, nil))
	defer server.Close()

	// test
	testRequestWithJSON(tst, server, "GET", "", `{"authorized":false,"success":true}`)
	testRequestWithJSON(tst, server, "POST", "", `{"authorized":false,"success":true}`)
}

func TestJhandler04(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Jhandler04. GET and POST with Panic")

	// create test server
	server := httptest.NewServer(Jhandler(jhandler04, nil))
	defer server.Close()

	// test
	testRequestWithJSON(tst, server, "GET", "", `{"authorized":false,"success":false,"error":"jhandler04 wants to stop"}`)
	testRequestWithJSON(tst, server, "POST", "", `{"authorized":false,"success":false,"error":"jhandler04 wants to stop"}`)
}
