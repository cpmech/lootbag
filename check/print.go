// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import "fmt"

// verboseMode controls message level
var verboseMode = false

// Verbose is an auxiliary function to set verbose mode
func Verbose() {
	verboseMode = true
}

// testTitle prints title of test
func testTitle(title string) {
	if verboseMode {
		fmt.Printf("\n=== %s =================\n", title)
		return
	}
	fmt.Printf("   . . . testing . . .   %s\n", title)
}
