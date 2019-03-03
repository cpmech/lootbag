// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/cpmech/lootbag/check"
)

func TestWriteFile01(tst *testing.T) {

	Verbose()
	TestTitle("WriteFile01. Write file")

	b0 := new(bytes.Buffer)
	b1 := new(bytes.Buffer)
	Ff(b0, "Hello World!\n")
	Ff(b1, "(using lootbag.lio.WriteFile)\n")

	dir := "/tmp"
	fn := "t_fileio_test_WriteFile01.txt"
	WriteFile(dir, fn, false, b0.Bytes(), b1.Bytes())

	res := ReadFile(filepath.Join(dir, fn))
	check.String(tst, "content of file", string(res), "Hello World!\n(using lootbag.lio.WriteFile)\n")
}
