// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// String checks string
func String(tst *testing.T, msg, a, b string) {
	if a != b {
		tst.Errorf("\n%q\nIS NOT EQUAL TO\n%q\n", a, b)
		return
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}

// Int64 checks int64
func Int64(tst *testing.T, msg string, a, b int64) {
	if a != b {
		tst.Errorf("%v != %v\n", a, b)
		return
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}

// Int checks int
func Int(tst *testing.T, msg string, a, b int) {
	if a != b {
		tst.Errorf("%v != %v\n", a, b)
		return
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}

// Float64 checks float64
func Float64(tst *testing.T, msg string, tol, a, b float64) {
	if math.IsNaN(a) {
		tst.Errorf("a=%v is NaN\n", a)
		return
	}
	if math.IsInf(a, 0) {
		tst.Errorf("a=%v is Inf\n", a)
		return
	}
	if math.IsNaN(b) {
		tst.Errorf("b=%v is NaN\n", b)
		return
	}
	if math.IsInf(b, 0) {
		tst.Errorf("b=%v is Inf\n", b)
		return
	}
	if math.Abs(a-b) > tol {
		tst.Errorf("%v != %v\n", a, b)
		return
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}

// Bools checks slice of bool
func Bools(tst *testing.T, msg string, a, b []bool) {
	if len(a) != len(b) {
		tst.Errorf("len(a)=%d != len(b)=%d\n", len(a), len(b))
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			tst.Errorf("a[%d]=%v != b[%d]=%v\n", i, a[i], i, b[i])
			return
		}
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}

// Time checks time.Time
func Time(tst *testing.T, msg string, a, b time.Time) {
	if a != b {
		tst.Errorf("%v != %v\n", a, b)
		return
	}
	if verboseMode {
		fmt.Printf("%s: OK\n", msg)
	}
}
