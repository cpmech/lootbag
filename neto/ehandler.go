// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"net/http"

	"github.com/cpmech/lootbag/lio"
)

// Ehandler makes a handler wrapping another handler in order to catch panics
func Ehandler(hfcn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() { // catch errors
			if err := recover(); err != nil {
				lio.Ff(w, RjsonFailed(err))
			}
		}()
		hfcn(w, r) // call the wrapped handler
	}
}

// EhandlerMW makes a handler wrapping another handler in order to catch panics
// (middleware version)
func EhandlerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { // catch errors
			if err := recover(); err != nil {
				lio.Ff(w, RjsonFailed(err))
			}
		}()
		next.ServeHTTP(w, r) // call the wrapped handler
	})
}
