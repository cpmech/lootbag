// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"net/http"
	"strings"

	"github.com/cpmech/lootbag/check"
	"github.com/go-chi/chi"
)

// SetFileServerRoute sets a route connected to a file server
// Based on https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//
//   Example:
//     workDir, _ := os.Getwd()
//     neto.SetFileServerRoute(router, "/", http.Dir(filepath.Join(workDir, "client/build")))
//
func SetFileServerRoute(router chi.Router, path string, root http.FileSystem) {

	// check path string
	if strings.ContainsAny(path, "{}*") {
		check.Panic("cannot have URL parameters in path string")
	}

	// Go fileserver
	fs := http.StripPrefix(path, http.FileServer(root))

	// fix path string
	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	// set route
	router.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
