// Example: If/Else

package main

import "fmt"

func main() {
	if 5%2 == 0 { // if statement: if 5/2 and remainder is 0, then
		fmt.Println("5 is even") // then execute this
	} else {
		fmt.Println("5 is odd") // if reminder is other than 0, then execute this
	}

	if 100%50 == 0 { // you can also use single if, as statement,
		// or if condition doesn't match then it will doesn't give anything
		fmt.Println("100 is divided by 50")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 { // you can also use, nested if/else
		fmt.Println(num, "have single digit")
	} else {
		fmt.Println(num, "have multiple digits")
	}
}
