package sampler

// Rand is a random number generator.
type Rand interface {
	Intn(n int) int
}

func perm[E any](rnd Rand, m []int) {
	// adopted from math/rand/perm.go
	// Copyright 2009 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE_go file.

	for i := 0; i < len(m); i++ {
		j := rnd.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
}

func shuffle(rnd Rand, n int, swap func(i, j int)) {
	// adopted from math/rand/perm.go
	// Copyright 2009 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE_go file.

	for i := n - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		swap(i, j)
	}
}

func fullShuffle[E any](rnd Rand, dst, items []E) []E {
	dst = append(dst, items...)
	p := dst[len(dst)-len(items):]
	shuffle(rnd, len(p), func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})

	return dst
}
