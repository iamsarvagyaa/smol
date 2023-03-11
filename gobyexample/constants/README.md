Constants

```go

// Example: Constants

package main

import ( // if we have many pkgs to import, we have to import like this
	"fmt"
	"math"
)

const s string = "constant-string" // declaring global constants

func main() {
	const n = 50505055 // declaring func block specific const
	const d = 3e20 / n // we can also declare const with expression

	fmt.Println(s)
	fmt.Println(d)
	fmt.Println(int64(n))    // a numeric constant don't have a type until we give
	fmt.Println(math.Sin(n)) // applying maths over const (n)
}

```