// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
)

// CallerInfo returns the file and line positions where an error occurred
//  idx -- use idx=2 to get the caller of Panic
func CallerInfo(idx int) {
	pc, file, line, ok := runtime.Caller(idx)
	if !ok {
		file, line = "?", 0
	}
	var fname string
	f := runtime.FuncForPC(pc)
	if f != nil {
		fname = f.Name()
	}
	if verboseMode {
		log.Printf("file = %s:%d\n", file, line)
		log.Printf("func = %s\n", fname)
	}
}

// Panic calls CallerInfo and panicks
func Panic(msg string, prm ...interface{}) {
	CallerInfo(4)
	CallerInfo(3)
	CallerInfo(2)
	panic(fmt.Sprintf(msg, prm...))
}

// MaybePanic calls Panic if doPanic is true; otherwise it does nothing
func MaybePanic(doPanic bool, msg string, prm ...interface{}) {
	if doPanic {
		Panic(msg, prm...)
	}
}

// Recover catches panics and call os.Exit(1) on 'panic'
func Recover() {
	if err := recover(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}

// RecoverTst catches panics in tests. Test will fail on 'panic'
func RecoverTst(tst *testing.T) {
	if err := recover(); err != nil {
		tst.Errorf("%v\n", err)
		tst.FailNow()
	}
}

// RecoverTstPanicIsOK catches panics in tests. Test must 'panic' to be OK
func RecoverTstPanicIsOK(tst *testing.T) {
	if err := recover(); err == nil {
		tst.Errorf("Test should have panicked\n")
		tst.FailNow()
	}
}
