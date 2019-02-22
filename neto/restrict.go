// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"net/http"
	"strings"

	"github.com/cpmech/lootbag/lio"
)

// Restricted is a http.Handler that checks (user)ID before handling requests
type Restricted struct {
	GiveAccess  func(r *http.Request) bool // [to be overwritten] returns "true" to give access; or "false" otherwise
	FailHandler http.HandlerFunc           // [optional] handles failed ID verifications
}

// Handler is a handler (middleware) that allows only ID matches to pass through
// NOTE: If FailHandler is nil, an ErrJSON string is returned
func (o *Restricted) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if o.GiveAccess(r) {
			next.ServeHTTP(w, r)
		} else {
			if o.FailHandler == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json")
				lio.Ff(w, RjsonUnauthorized())
			} else {
				o.FailHandler(w, r)
			}
		}
	})
}

// ExtractAuthorizationToken returns the 'Basic' or 'Bearer' authorization token
// given in the header of the request. Returns "" (empty) if not found
func ExtractAuthorizationToken(r *http.Request) (token string) {

	// bail on request without authorization header
	authHeader := r.Header.Get("authorization")
	if authHeader == "" {
		return // unauthorized
	}

	// bail on wrong number of words in authorization header
	bearerAndToken := strings.Split(authHeader, " ")
	if len(bearerAndToken) != 2 {
		return // unauthorized
	}

	// extract token
	token = bearerAndToken[1] // 0:Bearer, 1:Token
	return
}
