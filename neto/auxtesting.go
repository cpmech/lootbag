// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"net/http"
	"testing"

	"github.com/cpmech/lootbag/check"
)

// CheckGET checks if GET request works
// Returns response that may be <nil> in case of failure
func CheckGET(tst *testing.T, url string) (response *http.Response) {
	response, err := http.Get(url)
	if err != nil {
		tst.Errorf("GET failed: %v\n", err)
		return
	}
	return
}

// CheckResponse checks whether response is equal to correctBody or not
// NOTE: (1) reponse may be nil
//       (2) the response.Body data will be deleted/consumed by io.Copy()
func CheckResponse(tst *testing.T, response *http.Response, correctBody string) {
	if response == nil {
		tst.Errorf("cannot extract response.Body because response is <nil>\n")
		return
	}
	bodyText := ExtractResponseBodyText(response)
	check.String(tst, "response", bodyText, correctBody)
}
