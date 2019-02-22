// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"math"
	"testing"
	"time"
)

func TestCheck01(tst *testing.T) {

	// Verbose()
	testTitle("Check01. String")

	t := new(testing.T)
	String(t, "hello != yellow", "hello", "yellow")
	if !t.Failed() {
		tst.Errorf("test should have failed")
	}

	r := new(testing.T)
	String(r, "hello == hello", "hello", "hello")
	if r.Failed() {
		tst.Errorf("test should not have failed")
	}
}

func TestCheck02(tst *testing.T) {

	// Verbose()
	testTitle("Check02. Int64")

	t := new(testing.T)
	Int64(t, "123 != 456", 123, 456)
	if !t.Failed() {
		tst.Errorf("test should have failed")
	}

	r := new(testing.T)
	Int64(r, "123 == 123", 123, 123)
	if r.Failed() {
		tst.Errorf("test should not have failed")
	}
}

func TestCheck03(tst *testing.T) {

	// Verbose()
	testTitle("Check03. Int")

	t := new(testing.T)
	Int(t, "123 != 456", 123, 456)
	if !t.Failed() {
		tst.Errorf("test should have failed")
	}

	r := new(testing.T)
	Int(r, "123 != 123", 123, 123)
	if r.Failed() {
		tst.Errorf("test should not have failed")
	}
}

func TestCheck04(tst *testing.T) {

	// Verbose()
	testTitle("Check04. Float64")

	hitol := 1e-10
	lotol := 1e-8

	t := new(testing.T)
	Float64(t, "123.456 != 123.456000001", hitol, 123.456, 123.456000001)
	if !t.Failed() {
		tst.Errorf("(t) test should have failed")
	}

	r := new(testing.T)
	Float64(r, "123.456 == 123.456000001", lotol, 123.456, 123.456000001)
	if r.Failed() {
		tst.Errorf("(r) test should not have failed")
	}

	a := new(testing.T)
	Float64(a, "NaN != 123", hitol, math.NaN(), 123)
	if !a.Failed() {
		tst.Errorf("(a) test should have failed")
	}

	b := new(testing.T)
	Float64(b, "123 != NaN", hitol, 123, math.NaN())
	if !b.Failed() {
		tst.Errorf("(b) test should have failed")
	}

	c := new(testing.T)
	Float64(c, "Inf != 123", hitol, math.Inf(1), 123)
	if !c.Failed() {
		tst.Errorf("(c) test should have failed")
	}

	d := new(testing.T)
	Float64(d, "123 != Inf", hitol, 123, math.Inf(1))
	if !d.Failed() {
		tst.Errorf("(d) test should have failed")
	}
}

func TestCheck05(tst *testing.T) {

	// Verbose()
	testTitle("Check05. Bools")

	t := new(testing.T)
	Bools(t, "[true,true,false] != [true,false,false]", []bool{true, true, false}, []bool{true, false, false})
	if !t.Failed() {
		tst.Errorf("test should have failed")
	}

	r := new(testing.T)
	Bools(r, "[true,true,false] != [true,true,false]", []bool{true, true, false}, []bool{true, true, false})
	if r.Failed() {
		tst.Errorf("test should not have failed")
	}

	a := new(testing.T)
	Bools(a, "[true,true,false] != []", []bool{true, true, false}, nil)
	if !a.Failed() {
		tst.Errorf("test should have failed")
	}
}

func TestCheck06(tst *testing.T) {

	// Verbose()
	testTitle("Check06. Time")

	delta, _ := time.ParseDuration("100ms")
	now := time.Now()
	later := now.Add(delta)

	t := new(testing.T)
	Time(t, "now != later", now, later)
	if !t.Failed() {
		tst.Errorf("test should have failed")
	}

	r := new(testing.T)
	Time(r, "now == now", now, now)
	if r.Failed() {
		tst.Errorf("test should not have failed")
	}
}
