// Copyright © 2021 Jeffrey H. Johnson <trnsz@pobox.com>.
// Copyright © 2021 Gridfinity, LLC.
// Copyright © 2020 The Go Authors.
//
// All rights reserved.
//
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

// +build leaktest

package leaktestfe

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
