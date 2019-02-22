// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"testing"
)

func TestPanic01(tst *testing.T) {

	// Verbose()
	testTitle("Panic01. test must not panic")

	defer RecoverTst(tst)
}

func TestPanic02(tst *testing.T) {

	// Verbose()
	testTitle("Panic02. test must panic")

	defer RecoverTstPanicIsOK(tst)
	Panic("hello world: %v", "123")
}

func TestPanic03(tst *testing.T) {

	// Verbose()
	testTitle("Panic03. test will panic")

	defer RecoverTstPanicIsOK(tst)
	MaybePanic(true, "hello world: %v", "123")
}

func TestPanic04(tst *testing.T) {

	// Verbose()
	testTitle("Panic04. test will no panic")

	defer RecoverTst(tst)
	MaybePanic(false, "hello world: %v", "123")
}
