// Copyright 2020 Jeffrey H. Johnson.
// Copyright 2020 Gridfinity, LLC.
// Copyright 2019 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

// +build leaktest

package leaktestfe // import "go.gridfinity.dev/leaktestfe"

import (
	"testing"

	ltfe "go.uber.org/goleak"
)

func Leakplug(
	t *testing.T,
) {
	t.Log(
		"leaktestfe: enabled",
	)
	ltfe.VerifyNone(
		t,
	)
}
