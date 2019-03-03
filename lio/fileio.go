// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cpmech/lootbag/check"
)

// ReadFile reads bytes from a file
func ReadFile(fn string) (b []byte) {
	b, err := ioutil.ReadFile(os.ExpandEnv(fn))
	if err != nil {
		check.Panic("%v\n", err)
	}
	return
}

// WriteFile writes data to a new file
// dirout: directory for output. use "" or "." for the local dir
func WriteFile(dirout, fn string, verbose bool, data ...[]byte) {
	if dirout != "" && dirout != "." {
		os.MkdirAll(dirout, 0777)
		fn = filepath.Join(dirout, fn)
	}
	fil, err := os.Create(os.ExpandEnv(fn))
	if err != nil {
		check.Panic("cannot create file <%s>", fn)
	}
	defer fil.Close()
	for k := range data {
		if len(data[k]) > 0 {
			fil.Write(data[k])
		}
	}
	if verbose {
		Pf("file <%s> written\n", fn)
	}
}
