// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"fmt"
	"net/http"
)

// RedirectToHTTPS redirects HTTP to HTTPS
// Works with mux.Use command; e.g. mux.NewRouter().Use(wutil.RedirectToHTTPSRouter)
func RedirectToHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
