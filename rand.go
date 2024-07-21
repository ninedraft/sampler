package sampler

// Rand is a random number generator.
type Rand interface {
	IntN(n int) int
}

func perm[E any](rnd Rand, m []int) {
	// adopted from math/rand/perm.go
	// Copyright 2009 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE_go file.

	for i := 0; i < len(m); i++ {
		j := rnd.IntN(i + 1)
		m[i] = m[j]
		m[j] = i
	}
}

func fullShuffle[E any](rnd Rand, dst, items []E) []E {
	dst = append(dst, items...)
	p := dst[len(dst)-len(items):]

	for i := len(p) - 1; i > 0; i-- {
		j := rnd.IntN(i + 1)

		p[i], p[j] = p[j], p[i]
	}

	return dst
}
