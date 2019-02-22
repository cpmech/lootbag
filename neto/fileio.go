// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"html/template"

	"github.com/cpmech/lootbag/lio"
)

// ReadHTML reads a .html file into a template
func ReadHTML(filename string) *template.Template {
	htm := string(lio.ReadFile(filename))
	return template.Must(template.New(filename).Parse(htm))
}
