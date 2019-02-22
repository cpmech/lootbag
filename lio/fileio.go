// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"io/ioutil"
	"os"

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
