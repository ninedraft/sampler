// #nosec G404 // ignore weak random number generator in test code
package sampler_test

import (
	"math/rand"
	"testing"

	"github.com/ninedraft/sampler"
	"golang.org/x/exp/slices"
)

func BenchmarkAppend(b *testing.B) {
	rnd := rand.New(rand.NewSource(100))
	items := make([]int, 1000)
	dst := make([]int, 0, len(items))
	n := len(items) / 2
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst = dst[:0]
		dst = sampler.SampleAppend(rnd, dst, items, n)
	}
}

func TestSampleAppend(t *testing.T) {
	const seed = 100

	t.Run("n is 0", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))
		dst := []int{1, 2, 3}
		items := []int{4, 5, 6}

		got := sampler.SampleAppend(rnd, dst, items, 0)

		assertEqual(t, got, dst)
	})

	t.Run("dst is nil, n is full shuffle", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))
		var dst []int
		items := []int{1, 2, 3}

		got := sampler.SampleAppend(rnd, dst, items, -1)

		assertUnique(t, got)
		assertContainsAll(t, got, items)
	})

	t.Run("dst is nil, n is partial shuffle", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))
		var dst []int
		items := []int{1, 2, 3}

		got := sampler.SampleAppend(rnd, dst, items, 2)

		assertUnique(t, got)
		assertIntersection(t, got, items)
	})

	t.Run("dst is not nil, n is full shuffle", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))
		dst := []int{4, 5, 6}
		items := []int{1, 2, 3}

		got := sampler.SampleAppend(rnd, dst, items, -1)

		assertUnique(t, got)
		assertPrefix(t, got, dst)
		assertContainsAll(t, got, items)
	})

	t.Run("n is greater than or equal to len(items)", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))

		dst := []int{1, 2, 3}
		items := []int{4, 5, 6}

		got := sampler.SampleAppend(rnd, dst, items, 6)

		assertUnique(t, got)
		assertPrefix(t, got, dst)
		assertContainsAll(t, got, items)
	})

	t.Run("3*n is greater than len(items)", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))

		dst := []int{1, 2, 3}
		items := []int{4, 5, 6, 7, 8}

		got := sampler.SampleAppend(rnd, dst, items, 2)

		assertUnique(t, got)
		assertPrefix(t, got, dst)
		assertIntersection(t, got, items)
	})

	t.Run("n is less than len(items)", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewSource(seed))
		dst := []int{1, 2, 3}
		items := []int{4, 5, 6, 7, 8, 9}

		got := sampler.SampleAppend(rnd, dst, items, 2)

		assertUnique(t, got)
		assertPrefix(t, got, dst)
		assertIntersection(t, got, items)
	})
}

func assertEqual[E comparable](t *testing.T, got, want []E) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("want %v", want)
		t.Errorf(" got %v", got)
	}
}

func assertUnique[E comparable](t *testing.T, items []E) {
	t.Helper()

	set := make(map[E]struct{})
	for _, item := range items {
		if _, ok := set[item]; ok {
			t.Errorf("duplicate item %v", item)
		}
		set[item] = struct{}{}
	}
}

func assertContainsAll[E comparable](t *testing.T, items, want []E) {
	t.Helper()

	for _, item := range want {
		if !slices.Contains(items, item) {
			t.Errorf("item %v not found", item)
		}
	}
}

func assertIntersection[E comparable](t *testing.T, items, want []E) {
	t.Helper()

	for _, item := range items {
		if slices.Contains(want, item) {
			return // found
		}
	}

	t.Errorf("no intersection found between %v and %v", items, want)
}

func assertPrefix[E comparable](t *testing.T, items, want []E) {
	t.Helper()

	if len(items) < len(want) {
		t.Errorf("want %v, got %v", want, items)
		return
	}

	head := items[:len(want)]
	if !slices.Equal(head, want) {
		t.Error("prefix mismatch")
		t.Errorf("\twant %v", want)
		t.Errorf("\t got %v", head)
	}
}
