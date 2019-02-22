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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	lio.Pf("\n**************************************** HERE ****************************************\n\n")
	lio.Ff(w, "hello")
}

func TestRedirectToHTTPS(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("RedirectToHTTPS.")

	// create test server
	server := httptest.NewServer(RedirectToHTTPS(http.HandlerFunc(helloHandler)))
	defer server.Close()

	// test
	idToken := "123a567B"
	testRestricted(tst, server, "GET", idToken, "hello", false)
	testRestricted(tst, server, "POST", idToken, "hello", false)
}
