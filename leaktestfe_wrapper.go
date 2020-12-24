// Copyright 2020 Jeffrey H. Johnson.
// Copyright 2020 Gridfinity, LLC.
// Copyright 2019 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

// +build !leaktest,leaktest

package leaktestfe // import "go.gridfinity.dev/leaktestfe"

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
