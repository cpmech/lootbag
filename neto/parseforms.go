// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cpmech/lootbag/check"
)

// FormGetParam returns parameter in form
//   Input:
//     paramName -- name of parameter in form
//     r         -- request
//     parseForm -- do call r.ParseForm or r.ParseMultipartForm
//     multiPart -- call r.ParseMultiForm instead of r.ParseForm
//     doPanic   -- panic on errors; otherwise return silently
//   Output:
//     paramValue -- value of parameter
func FormGetParam(paramName string, r *http.Request, parseForm, multiPart, doPanic bool) (paramValue string) {

	// parse form
	var err error
	if parseForm {
		if multiPart {
			err = r.ParseMultipartForm(32 << 20)
		} else {
			err = r.ParseForm()
		}
		if err != nil {
			check.MaybePanic(doPanic, "cannot parse form: %v\n", err)
			return
		}
	}

	// get parameter
	formData, ok := r.Form[paramName]
	if !ok {
		check.MaybePanic(doPanic, "cannot extract parameter named %q in form\n", paramName)
		return
	}
	if len(formData) != 1 {
		check.MaybePanic(doPanic, "number of parameters in form is incorrect. %d != 1\n", len(formData))
		return
	}

	// results
	paramValue = formData[0]
	return
}

// FormGetFile gets file from form
//   Input:
//     paramName   -- file field name in form; e.g. "image"
//     r           -- request
//     parseForm   -- do call r.ParseMultipartForm
//     doPanic     -- panic on errors; otherwise return silently
//   Output:
//     file -- io.Reader that can be use to save file as in:
//             io.Copy(destination, file)
//
//     NOTE: remember to call file.Close() after using its contents
//
func FormGetFile(paramName string, r *http.Request, parseForm, doPanic bool) (filename string, file io.Reader) {

	// parse form
	var err error
	if parseForm {
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			check.MaybePanic(doPanic, "cannot parse form: %v\n", err)
			return
		}
	}

	// get file
	f, handle, err := r.FormFile(paramName)
	if err != nil {
		check.MaybePanic(doPanic, "cannot get file from form: %v\n", err)
		return
	}
	return filepath.Base(handle.Filename), f
}

// FormGetAndSaveFile gets file from form and save into destination
//   Input:
//     dirout      -- where to save file. NOTE: new directory will be created
//     paramName   -- file field name in form; e.g. "image"
//     r           -- request
//     parseForm   -- do call r.ParseMultipartForm
//     doPanic     -- panic on errors; otherwise return silently
//   Output:
//     path -- location of file on success; returns "" if failed
func FormGetAndSaveFile(dirout, paramName string, r *http.Request, parseForm, doPanic bool) (path string) {

	// parse form
	var err error
	if parseForm {
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			check.MaybePanic(doPanic, "cannot parse form: %v\n", err)
			return
		}
	}

	// get file
	file, handler, err := r.FormFile(paramName)
	if err != nil {
		check.MaybePanic(doPanic, "cannot get file from form: %v\n", err)
		return
	}
	defer file.Close()

	// save file
	fn := filepath.Base(handler.Filename)
	fp := filepath.Join(dirout, fn)
	os.MkdirAll(dirout, 0777)
	f, err := os.Create(fp)
	if err != nil {
		check.MaybePanic(doPanic, "cannot save file: %v\n", err)
		return
	}
	defer f.Close()

	// copy data
	_, err = io.Copy(f, file)
	if err != nil {
		check.MaybePanic(doPanic, "cannot copy data: %v\n", err)
		return
	}

	// success
	path = fp
	return
}
