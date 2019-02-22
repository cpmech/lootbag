// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

// testRestricted runs test that makes a GET or POST request with authentication credentials
func testRestricted(tst *testing.T, server *httptest.Server, method, idToken, correctOutput string, unauthorizedIsOk bool) {

	// request
	var request *http.Request
	var err error
	switch method {
	case "GET":
		request, err = http.NewRequest(method, server.URL, nil)

	case "POST":
		body := bytes.NewBuffer([]byte(`{"somenumber":123}`))
		request, err = http.NewRequest(method, server.URL, body)
	}
	if err != nil {
		tst.Errorf("cannot create request: %v\n", err)
		return
	}

	// header of request
	request.Header.Add("Authorization", "Bearer "+idToken)
	if method == "POST" {
		request.Header.Set("Content-Type", "application/json")
	}

	// make request
	response, err := server.Client().Do(request)
	if err != nil {
		tst.Errorf("cannot make request: %v\n", err)
		return
	}
	defer func() {
		if response != nil {
			response.Body.Close()
		}
	}()

	// check response
	correctCode := 200
	if unauthorizedIsOk {
		correctCode = 401
	}
	if response.StatusCode != correctCode {
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

func TestRestrict01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Restrict01.")

	// restrict handler
	restrict := &Restricted{
		GiveAccess: func(r *http.Request) bool {
			ID := ExtractAuthorizationToken(r)
			return ID == "123a567B"
		},
	}

	// next handler
	authOK := RjsonAllGood()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lio.Ff(w, authOK)
	})

	// create test server
	server := httptest.NewServer(restrict.Handler(next))
	defer server.Close()

	// test with good idToken
	idToken := "123a567B"
	testRestricted(tst, server, "GET", idToken, authOK, false)
	testRestricted(tst, server, "POST", idToken, authOK, false)

	// test with bad idToken
	idToken = "12345678"
	authFAIL := RjsonUnauthorized()
	testRestricted(tst, server, "GET", idToken, authFAIL, true)
	testRestricted(tst, server, "POST", idToken, authFAIL, true)
}
