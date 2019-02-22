// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

func TestResults01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Results01. From JSON")

	res := NewResultsFromJSON([]byte(`{"authorized":true,"success":true,"data":{"memberA":"A","memberB":123}}`))
	if !res.Authorized {
		tst.Errorf("authorized should be true\n")
		return
	}
	if !res.Success {
		tst.Errorf("success should be true\n")
		return
	}
	if iface, ok := res.Data["memberA"]; ok {
		check.String(tst, "memberA", iface.(string), "A")
	} else {
		tst.Errorf("memberA should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberB"]; ok {
		check.Float64(tst, "memberB", 1e-15, iface.(float64), 123)
	} else {
		tst.Errorf("memberB should be present as interface{}\n")
		return
	}
}

func TestResults02(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Results02. All good with data")

	jres := RjsonAllGoodWithData("memberA", "A", "memberB", 123, "memberC", false, "memberD", 123.456)
	check.String(tst, "jres", jres, `{"authorized":true,"success":true,"data":{"memberA":"A","memberB":123,"memberC":false,"memberD":123.456}}`)

	res := NewResultsFromJSON([]byte(jres))
	if !res.Authorized {
		tst.Errorf("authorized should be true\n")
		return
	}
	if !res.Success {
		tst.Errorf("success should be true\n")
		return
	}
	if iface, ok := res.Data["memberA"]; ok {
		check.String(tst, "memberA", iface.(string), "A")
	} else {
		tst.Errorf("memberA should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberB"]; ok {
		check.Float64(tst, "memberB", 1e-15, iface.(float64), 123)
	} else {
		tst.Errorf("memberB should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberC"]; ok {
		if iface.(bool) {
			tst.Errorf("memberC should be false\n")
			return
		}
		lio.Pf("memberC: OK\n")
	} else {
		tst.Errorf("memberB should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberD"]; ok {
		check.Float64(tst, "memberD", 1e-15, iface.(float64), 123.456)
	} else {
		tst.Errorf("memberD should be present as interface{}\n")
		return
	}
}

func TestResults03(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("Results03. All good with data")

	jres := RjsonAllGoodWithDataString(`{"memberA":"A","memberB":123,"memberC":false,"memberD":123.456}`)
	check.String(tst, "jres", jres, `{"authorized":true,"success":true,"data":{"memberA":"A","memberB":123,"memberC":false,"memberD":123.456}}`)

	res := NewResultsFromJSON([]byte(jres))
	if !res.Authorized {
		tst.Errorf("authorized should be true\n")
		return
	}
	if !res.Success {
		tst.Errorf("success should be true\n")
		return
	}
	if iface, ok := res.Data["memberA"]; ok {
		check.String(tst, "memberA", iface.(string), "A")
	} else {
		tst.Errorf("memberA should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberB"]; ok {
		check.Float64(tst, "memberB", 1e-15, iface.(float64), 123)
	} else {
		tst.Errorf("memberB should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberC"]; ok {
		if iface.(bool) {
			tst.Errorf("memberC should be false\n")
			return
		}
		lio.Pf("memberC: OK\n")
	} else {
		tst.Errorf("memberB should be present as interface{}\n")
		return
	}
	if iface, ok := res.Data["memberD"]; ok {
		check.Float64(tst, "memberD", 1e-15, iface.(float64), 123.456)
	} else {
		tst.Errorf("memberD should be present as interface{}\n")
		return
	}
}
