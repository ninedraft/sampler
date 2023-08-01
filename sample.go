package sampler

import (
	"sync"

	"golang.org/x/exp/slices"
)

// Sample returns a random selection of n items from items.
// Items with unique indexes are selected.
//
//	[1 2 3 4 5] -> always unique items in result
//	[1 2 3 4 4] -> 4 can be selected twice
func Sample[E any](rnd Rand, items []E, n int) []E {
	return SampleAppend(rnd, nil, items, n)
}

// SampleAppend returns a random selection of n items from items.
// Items are appended to dst, which is grown if necessary.
// Items with unique indexes are selected.
//
//	[1 2 3 4 5] -> always unique items in result
//	[1 2 3 4 4] -> 4 can be selected twice
func SampleAppend[E any](rnd Rand, dst, items []E, n int) []E {
	if n == 0 {
		return dst
	}

	if n < 0 || n >= len(items) {
		return fullShuffle(rnd, dst, items)
	}

	if dst == nil {
		dst = make([]E, 0, n)
	}

	if 3*n > len(items) {
		return permutate(rnd, dst, items, n)
	}

	set := make(map[int]struct{}, n)
	for len(set) < n {
		i := rnd.Intn(len(items))
		if _, ok := set[i]; !ok {
			set[i] = struct{}{}
			dst = append(dst, items[i])
		}
	}

	return dst
}

func permutate[E any](rnd Rand, dst, items []E, n int) []E {
	ii := getIndexes(n)
	defer releaseIndexes(ii)

	perm[int](rnd, ii.values)

	for _, i := range ii.values {
		dst = append(dst, items[i])
	}
	return dst
}

type indexes struct {
	values []int
}

var indexesPool = &sync.Pool{
	New: func() any {
		return &indexes{
			values: make([]int, 64),
		}
	},
}

func getIndexes(n int) *indexes {
	//nolint:errcheck // ignore casting error - we know that we have *indexes
	ii := indexesPool.Get().(*indexes)
	ii.values = slices.Grow(ii.values, n)[:n]
	return ii
}

func releaseIndexes(ii *indexes) {
	if len(ii.values) > 1024*1024 {
		return
	}
	ii.values = ii.values[:0]
	indexesPool.Put(ii)
}
