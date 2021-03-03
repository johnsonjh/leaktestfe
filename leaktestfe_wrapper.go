// Copyright © 2021 Jeffrey H. Johnson <trnsz@pobox.com>.
// Copyright © 2021 Gridfinity, LLC.
// Copyright © 2020 The Go Authors.
//
// All rights reserved.
//
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

// +build !leaktest,leaktest

package leaktestfe

import (
	"testing"
)

func VerifyNone(
	t *testing.T,
) {
	Leakplug(
		t,
	)
}
