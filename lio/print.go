// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"fmt"
	"io"

	"github.com/cpmech/lootbag/check"
)

// verboseMode controls message level
var verboseMode = false

// Verbose is an auxiliary function to set verbose mode
// NOTE: also sets check.Verbose() accordingly
func Verbose() {
	verboseMode = true
	check.Verbose()
}

// IsVerbose returns verbose mode status
func IsVerbose() bool {
	return verboseMode
}

// Pl prints a new line
func Pl() {
	if !verboseMode {
		return
	}
	fmt.Printf("\n")
}

// Pf prints formatted string
func Pf(msg string, prm ...interface{}) {
	if !verboseMode {
		return
	}
	fmt.Printf(msg, prm...)
}

// Sf wraps Sprintf
func Sf(msg string, prm ...interface{}) string {
	return fmt.Sprintf(msg, prm...)
}

// Ff wraps Sprintf
func Ff(w io.Writer, msg string, prm ...interface{}) {
	_, err := fmt.Fprintf(w, msg, prm...)
	if err != nil {
		panic(fmt.Sprintf("cannot write using Fprintf: %v\n", err))
	}
}

// TestTitle prints title of test
func TestTitle(title string) {
	if verboseMode {
		fmt.Printf("\n=== %s =================\n", title)
		return
	}
	fmt.Printf("   . . . testing . . .   %s\n", title)
}
