package sampler

// Choice returns a random item from items.
// If items is empty, Choice panics.
func Choice[E any](rnd Rand, items []E) E {
	if len(items) == 0 {
		panic("sample.Choice called with empty items")
	}

	return items[rnd.IntN(len(items))]
}

// Choices returns a random selection of n items from items.
// Items are not guaranteed to be unique.
func Choices[E any](rnd Rand, items []E, n int) []E {
	return ChoicesAppend(rnd, nil, items, n)
}

// ChoicesAppend returns a random selection of n items from items.
// Items are appended to dst, which is grown if necessary.
// Items are not guaranteed to be unique.
func ChoicesAppend[E any](rnd Rand, dst, items []E, n int) []E {
	if n == 0 {
		return dst
	}

	if n < 0 {
		n = len(items)
	}

	if dst == nil {
		dst = make([]E, 0, n)
	}

	for i := 0; i < n; i++ {
		idx := rnd.IntN(len(items))
		dst = append(dst, items[idx])
	}

	return dst
}
