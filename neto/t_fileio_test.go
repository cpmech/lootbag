// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"testing"

	"github.com/cpmech/lootbag/check"
	"github.com/cpmech/lootbag/lio"
)

func TestReadHTML01(tst *testing.T) {

	// lio.Verbose()
	lio.TestTitle("ReadHTML01.")

	htmlTmp := ReadHTML("samples/sample01.html")
	htmlBuf := bytes.NewBuffer(nil)
	htmlTmp.Execute(htmlBuf, nil)
	check.String(tst, "sample01.html", string(htmlBuf.Bytes()), `<!DOCTYPE html>
<html>
<body>
<h1>Heading</h1>
<p>Paragraph.</p>
</body>
</html>
`)
}
