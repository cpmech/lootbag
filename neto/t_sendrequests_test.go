// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

func TestSendRequest01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendRequest01. GET")

	// create test server
	server := httptest.NewServer(Ehandler(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tester"))
	}))
	defer server.Close()

	// test
	responseBody := SendGetRequest(server.URL)
	check.String(tst, "response", string(responseBody), "hello tester")
}

func TestSendRequest02(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendRequest02. GET (restricted)")

	// restrict handler
	restrict := &Restricted{
		GiveAccess: func(r *http.Request) bool {
			ID := ExtractAuthorizationToken(r)
			return ID == "123a567B"
		},
	}

	// create test server
	server := httptest.NewServer(restrict.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tester"))
	})))
	defer server.Close()

	// test
	responseBody := SendGetRequestWithAuth(server.URL, "123a567B")
	check.String(tst, "response", string(responseBody), "hello tester")
}

func TestSendRequest03(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendRequest03. POST with JSON")

	// create test server
	server := httptest.NewServer(Jhandler(jhandler01, jparser))
	defer server.Close()

	// correct output
	inputJSON := `{"id":123, "name":"dorival"}`
	correctOutput := `{"authorized":false,"gotID":123,"gotName":"dorival","success":true}`

	// test
	responseBody := SendPostRequest(server.URL, inputJSON)
	check.String(tst, "response", string(responseBody), correctOutput)
}

func TestSendRequest04(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendRequest04. POST with JSON (restricted)")

	// restrict handler
	restrict := &Restricted{
		GiveAccess: func(r *http.Request) bool {
			ID := ExtractAuthorizationToken(r)
			return ID == "123a567B"
		},
	}

	// create test server
	server := httptest.NewServer(restrict.Handler(Jhandler(jhandler01, jparser)))
	defer server.Close()

	// correct output
	inputJSON := `{"id":123, "name":"dorival"}`
	correctOutput := `{"authorized":false,"gotID":123,"gotName":"dorival","success":true}`

	// test
	responseBody := SendPostRequestWithAuth(server.URL, []byte(inputJSON), "123a567B")
	check.String(tst, "response", string(responseBody), correctOutput)
}

func TestSendForm01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendForm01. Send form with image file (restricted)")

	// restrict handler
	restrict := &Restricted{
		GiveAccess: func(r *http.Request) bool {
			ID := ExtractAuthorizationToken(r)
			return ID == "123a567B"
		},
	}

	// create test server
	parseForm, multipart, doPanic := true, true, true
	server := httptest.NewServer(restrict.Handler(Ehandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		owner := FormGetParam("owner", r, parseForm, multipart, doPanic)
		title := FormGetParam("title", r, parseForm, multipart, doPanic)
		path := FormGetAndSaveFile("/tmp/neto/files", "image", r, parseForm, doPanic)
		WjsonAllGoodWithData(w, "owner", owner, "title", title, "path", path)
	}))))
	defer server.Close()

	// POST form
	responseBody := SendFormRequestWithAuth(server.URL, "123a567B", "image", true, false,
		"owner", "tester@testing.co",
		"title", "Just Testing",
		"image", "./samples/doc.png",
	)

	// check
	correctOutput := `{"authorized":true,"success":true,"data":{"owner":"tester@testing.co","title":"Just Testing","path":"/tmp/neto/files/doc.png"}}`
	check.String(tst, "response", string(responseBody), correctOutput)
}

func TestSendForm02(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("SendForm02. Send form with text file (restricted)")

	// restrict handler
	restrict := &Restricted{
		GiveAccess: func(r *http.Request) bool {
			ID := ExtractAuthorizationToken(r)
			return ID == "123a567B"
		},
	}

	// create test server
	parseForm, doPanic := true, true
	server := httptest.NewServer(restrict.Handler(Ehandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn, file := FormGetFile("message", r, parseForm, doPanic)
		content := new(bytes.Buffer)
		io.Copy(content, file)
		WjsonAllGoodWithData(w, "fn", fn, "content", content.String())
	}))))
	defer server.Close()

	// POST form
	responseBody := SendFormRequestWithAuth(server.URL, "123a567B", "message", true, false,
		"message", "./samples/hello.txt",
	)

	// check
	correctOutput := `{"authorized":true,"success":true,"data":{"fn":"hello.txt","content":"Hello World 123\n"}}`
	check.String(tst, "response", string(responseBody), correctOutput)
}
