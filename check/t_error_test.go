// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import "testing"

func TestError01(tst *testing.T) {

	// Verbose()
	testTitle("Error01. Err")

	e := Err("1-2-3-%s", "4-5-6")
	if e.Error() != "1-2-3-4-5-6" {
		tst.Errorf("strings should be equal")
	}
}
