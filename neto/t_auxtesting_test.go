// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpmech/lootbag/lio"
)

func TestCheckGET01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("CheckGET01. correct body and nil response")

	// handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		lio.Ff(w, "hello")
	}

	// create test server
	server := httptest.NewServer(Ehandler(handler))
	defer server.Close()

	// make request and check response
	response := CheckGET(tst, server.URL)
	CheckResponse(tst, response, "hello")

	// nil response
	t0 := new(testing.T)
	CheckResponse(t0, nil, "hello")
	if !t0.Failed() {
		tst.Errorf("CheckResponse should have failed due to nil response\n")
		return
	}

	// wrong url
	t1 := new(testing.T)
	CheckGET(t1, server.URL+"noise")
	if !t1.Failed() {
		tst.Errorf("CheckGET with wrong url should have failed\n")
		return
	}
}

func TestCheckGET02(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("CheckGET02. wrong body")

	// handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		lio.Ff(w, "hello")
	}

	// create test server
	server := httptest.NewServer(Ehandler(handler))
	defer server.Close()

	// make request
	t0 := new(testing.T)
	response := CheckGET(t0, server.URL)
	if t0.Failed() {
		tst.Errorf("CheckGET should NOT have failed\n")
		return
	}

	// check response
	t1 := new(testing.T)
	CheckResponse(t1, response, "dorival")
	if !t1.Failed() {
		tst.Errorf("CheckResponse should have failed\n")
		return
	}
}
