// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"testing"

	"github.com/cpmech/lootbag/lio"
)

func testDatastore01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Datastore01. emulator")

	stop := DatastoreEmulator("epop-web", "8081")
	defer stop()
}
