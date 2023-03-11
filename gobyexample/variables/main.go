// Example: Variables

package main

import "fmt"

func main() {
	var a = "golang" // declaring variable without values type, go will automatically select
	fmt.Println(a)

	var b, c int = 1, 2 // you can also declare multiple variables, and also values type here isinteger
	fmt.Println(b, c)

	var d = true // go can automatically detects the values type
	fmt.Println(d)

	var e int // zero-valued variable because, we haven't given any values of var `e`
	fmt.Println(e)

	f := "bruuu"
	fmt.Println(f) // `:=` is used for shorthand, we can also define variable like this

	g, h := "henlo", 69 // here we are declaring multiple variables w/ shorthand
	fmt.Println(g, h)
}
