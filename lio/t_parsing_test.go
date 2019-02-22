// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lio

import (
	"testing"

	"github.com/cpmech/lootbag/check"
)

func TestParsing01(tst *testing.T) {

	//Verbose()
	TestTitle("Parsing01. Atob, Atoi, Atof")

	if !Atob("1") {
		tst.Errorf("Atob(\"1\") should have returned true\n")
		return
	}
	if !Atob("true") {
		tst.Errorf("Atob(\"true\") should have returned true\n")
	}
	if Atob("0") {
		tst.Errorf("Atob(\"0\") should have returned false\n")
	}
	if Atob("false") {
		tst.Errorf("Atob(\"false\") should have returned false\n")
	}

	if Itob(0) {
		tst.Errorf("Itob(0) should have returned false\n")
	}
	if !Itob(-1) {
		tst.Errorf("Itob(-1) should have returned true\n")
	}
	if !Itob(+1) {
		tst.Errorf("Itob(+1) should have returned true\n")
	}

	check.Int(tst, "true => 1", Btoi(true), 1)
	check.Int(tst, "false => 0", Btoi(false), 0)

	check.Int(tst, "\"123\" => 123", Atoi("123"), 123)

	check.String(tst, "true => \"true\"", Btoa(true), "true")
	check.String(tst, "false => \"false\"", Btoa(false), "false")

	check.Float64(tst, "\"123.456\" => 123.456", 1e-15, Atof("123.456"), 123.456)
}

func TestParsing02(tst *testing.T) {

	//Verbose()
	TestTitle("Parsing02. Atob panic")

	defer check.RecoverTstPanicIsOK(tst)
	Atob("dorival")
}

func TestParsing03(tst *testing.T) {

	//Verbose()
	TestTitle("Parsing03. Atoi panic")

	defer check.RecoverTstPanicIsOK(tst)
	Atoi("dorival")
}

func TestParsing04(tst *testing.T) {

	//Verbose()
	TestTitle("Parsing04. Atof panic")

	defer check.RecoverTstPanicIsOK(tst)
	Atof("dorival")
}
