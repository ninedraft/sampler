[![Go Reference](https://pkg.go.dev/badge/github.com/ninedraft/sampler.svg)](https://pkg.go.dev/github.com/ninedraft/sampler)
# sampler

Sampler is a golang library for sampling from a slice of items.

## Install

```
go get github.com/ninedraft/sampler
```

## Example usage 

```go
import "math/rand/v2" 

var rnd *rand.Rand = ...
items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// acquire a random subset of $items
sample := sampler.Sample(rnd, items, 3)
fmt.Println(sample)
// Output: [8 1 3]


// acquire N random choises from $items
choices := sampler.Choices(rnd, items, 3)
fmt.Println(choices)
// Output: [2 1 1]

// acquire choice a single item from $items
choice := sampler.Choice(rnd, items)
fmt.Println(choice)
// Output: 5
```

## License
APACHE 2.0