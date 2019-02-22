// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import "fmt"

// Err returns a new error
func Err(msg string, prm ...interface{}) error {
	return fmt.Errorf(msg, prm...)
}
