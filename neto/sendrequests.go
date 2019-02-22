// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/cpmech/lootbag/check"
)

// SendGetRequest sends a GET request
func SendGetRequest(url string) (responseBody []byte) {

	// get request
	resp, err := http.Get(url)
	if err != nil {
		check.Panic("cannot send GET request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendGetRequestWithAuth sends a GET request with Authorization header
func SendGetRequestWithAuth(url, token string) (responseBody []byte) {

	// set request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		check.Panic("cannot create request: %v\n", err)
	}

	// set header
	request.Header.Set("Authorization", "Bearer "+token)

	// post request
	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		check.Panic("cannot send GET request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendPostRequest sends a POST request with JSON data
func SendPostRequest(url, data string) (responseBody []byte) {

	// post request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		check.Panic("cannot send POST request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendPostRequestWithAuth sends a POST request with JSON data and Authorization header
func SendPostRequestWithAuth(url string, data []byte, token string) (responseBody []byte) {

	// set request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		check.Panic("cannot create request: %v\n", err)
	}

	// set header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	// post request
	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		check.Panic("cannot send POST request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendFormRequestWithAuth sends a POST request with form data and Authorization header
//    url       -- url of POST request
//    token     -- authentication token
//    fileField -- file field name to consider pairs[0] as a file; e.g. "image"
//                 may be "" if no files are to be given in pairs
//    withFile  -- pairs has a file item
//    pairs     -- an even list of "param", value to build the form
func SendFormRequestWithAuth(url, token, fileField string, withFile, PUT bool, pairs ...string) (responseBody []byte) {

	// check number of pairs
	nargs := len(pairs)
	if nargs%2 != 0 {
		check.Panic("number of items in pairs must be even.\npairs = %q\n", pairs)
	}

	// set form fields
	form := new(bytes.Buffer)
	m := multipart.NewWriter(form)
	for i := 0; i < nargs; i += 2 {

		// key - value
		key := pairs[i]
		val := pairs[i+1]

		// file
		if key == fileField && withFile {
			fw, err := m.CreateFormFile(key, val)
			if err != nil {
				check.Panic("cannot create file writer: %v\n", err)
			}
			fh, err := os.Open(val)
			if err != nil {
				check.Panic("os.Open(%q) failed: %v\n", val, err)
			}
			defer fh.Close()
			_, err = io.Copy(fw, fh)
			if err != nil {
				check.Panic("copy failed: %v\n", err)
			}

			// normal field
		} else {
			m.WriteField(key, val)
		}
	}
	contentType := m.FormDataContentType()
	m.Close()

	// set request
	method := "POST"
	if PUT {
		method = "PUT"
	}
	request, err := http.NewRequest(method, url, form)
	if err != nil {
		check.Panic("cannot create request: %v\n", err)
	}

	// set header
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Authorization", "Bearer "+token)

	// post request
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		check.Panic("cannot send %s request: %v\n", method, err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendDeleteRequestWithAuth sends a DELETE request with Authorization header
func SendDeleteRequestWithAuth(url, token string) (responseBody []byte) {

	// set request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		check.Panic("cannot create request: %v\n", err)
	}

	// set header
	request.Header.Set("Authorization", "Bearer "+token)

	// post request
	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		check.Panic("cannot send DELETE request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}

// SendPutRequestWithAuth sends a PUT request with JSON data and Authorization header
func SendPutRequestWithAuth(url string, data []byte, token string) (responseBody []byte) {

	// set request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		check.Panic("cannot create request: %v\n", err)
	}

	// set header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	// post request
	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		check.Panic("cannot send POST request: %v\n", err)
	}
	defer resp.Body.Close()

	// check status
	if resp.StatusCode != http.StatusOK {
		check.Panic("unexpected http status code: %d\n", resp.StatusCode)
	}

	// extract response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		check.Panic("cannot read response body: %v\n", err)
	}
	return body
}
