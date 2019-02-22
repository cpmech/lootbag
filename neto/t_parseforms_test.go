// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

// testRequestWithSimpleForm runs test that makes a GET or POST request with form parameters
func testRequestWithSimpleForm(tst *testing.T, server *httptest.Server, method, paramName, paramValue, correctOutput string) {

	// url with parameter
	url := server.URL + "?" + paramName + "=" + paramValue

	// make request
	var response *http.Response
	var err error
	if method == "GET" {
		response, err = http.Get(url)
	} else {
		response, err = http.Post(url, "application/json", nil)
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

func paramHandler(w http.ResponseWriter, r *http.Request) {
	returnedValue := FormGetParam("parameterABC", r, true, false, false)
	lio.Ff(w, returnedValue)
}

func TestGetParam01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("GetParam01.")

	// create test server
	server := httptest.NewServer(http.HandlerFunc(paramHandler))
	defer server.Close()

	// test => ok
	testRequestWithSimpleForm(tst, server, "GET", "parameterABC", "value123", "value123")
	testRequestWithSimpleForm(tst, server, "POST", "parameterABC", "value123", "value123")

	// test => unavailable
	testRequestWithSimpleForm(tst, server, "GET", "parameter", "value123", "")
	testRequestWithSimpleForm(tst, server, "POST", "parameter", "value123", "")
}
