// Copyright 2021 Gridfinity, LLC.
// Copyright 2020 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

// +build leaktest !purego

package leaktestfe_test

import (
	"testing"

	ltfe "go.gridfinity.dev/leaktestfe"
)

// TestLeakPlugOn -> Enable
func TestLeakPlugOn(
	t *testing.T,
) {
	ltfe.Leakplug(
		t,
	)
}
