// Copyright © 2021 Jeffrey H. Johnson <trnsz@pobox.com>.
// Copyright © 2021 Gridfinity, LLC.
// Copyright © 2019 The Go Authors.
//
// All rights reserved.
//
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

//go:build !leaktest || purego
// +build !leaktest purego

package leaktestfe_test

import (
	"testing"

	ltfe "github.com/johnsonjh/leaktestfe"
)

// TestLeakPlugOff -> Disabled
func TestLeakPlugOff(
	t *testing.T,
) {
	t.Log(
		"leaktestfe: disabled",
	)
	ltfe.Leakplug(
		t,
	)
}
