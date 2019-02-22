// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"testing"
)

func TestSf01(tst *testing.T) {

	//Verbose()
	TestTitle("Sf01. String formatting")

	res := Sf("%04d", 123)
	if res != "0123" {
		tst.Errorf("res = %q. want = \"123\"\n", res)
	}
}
