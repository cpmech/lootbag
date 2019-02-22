// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cpmech/lootbag/check"
)

// ExtractResponseBodyText returns text from response.Body
func ExtractResponseBodyText(response *http.Response) (bodyText string) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, response.Body)
	if err != nil {
		check.Panic("cannot read response.Body. io.Copy failed: %v\n", err)
	}
	bodyText = string(buf.Bytes())
	return
}
