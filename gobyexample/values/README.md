Values

```go

// Example: Values

package main

import "fmt"

func main() {
	fmt.Println("string")                   // string
	fmt.Println("string" + "concatenation") // string concatenation
	fmt.Println("1" + "2")                  // this will concatenate both numbers
	fmt.Println(1 + 2)                      // integer
	fmt.Println("5 * 2 =", 5*2)             // strings & integers
	fmt.Println("7.0 / 2.0", 7.0/2.0)       // floats
	fmt.Println(true)                       // boolean
	fmt.Println(true || false)              // result: true, boolean w/ boolean operator
	fmt.Println(true && false)              // result: false, boolean w/ boolean operator
	fmt.Println(!false)                     // result: true, boolean w/ boolean operator
}

```