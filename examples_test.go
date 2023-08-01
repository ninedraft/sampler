// #nosec G404 // ignore weak random number generator in test code
package sampler_test

import (
	"fmt"
	"math/rand"

	"github.com/ninedraft/sampler"
)

func ExampleSample() {
	rnd := rand.New(rand.NewSource(100))

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sample := sampler.Sample(rnd, items, 3)

	fmt.Println(sample)
	// Output: [5 3 8]
}

func ExampleChoices() {
	rnd := rand.New(rand.NewSource(100))

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	choices := sampler.Choices(rnd, items, 3)

	fmt.Println(choices)
	// Output: [5 3 8]
}

func ExampleChoice() {
	rnd := rand.New(rand.NewSource(100))

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	choice := sampler.Choice(rnd, items)

	fmt.Println(choice)
	// Output: 5
}
