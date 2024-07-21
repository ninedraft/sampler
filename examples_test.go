// #nosec G404 // ignore weak random number generator in test code
package sampler_test

import (
	"fmt"

	"github.com/ninedraft/sampler"
)

func ExampleSample() {
	rnd := newFixedRnd()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sample := sampler.Sample(rnd, items, 3)

	fmt.Println(sample)
	// Output: [2 1 3]
}

func ExampleChoices() {
	rnd := newFixedRnd()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	choices := sampler.Choices(rnd, items, 3)

	fmt.Println(choices)
	// Output: [2 1 3]
}

func ExampleChoice() {
	rnd := newFixedRnd()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	choice := sampler.Choice(rnd, items)

	fmt.Println(choice)
	// Output: 2
}
