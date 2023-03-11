// Example: For

package main

import "fmt"

func main() {
	i := 1       // declaring `i` var
	for i <= 9 { // for loop, which loops till `i` is equal or greater than 9
		fmt.Println(i) // while looping over, print the value of `i`
		i = i + 1      // then exceed by 1 till `i=9`
	}

	for j := 6; j <= 18; j++ { // classic loop
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break // you can also break the loop, if you're condition doesn't relate
		// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // you can also continue to next loop, if you're condition doesn't relate
		}

		fmt.Println(n) // this will print odd numers because in above if condition there is continue statement, if n is completely divided by 0 then continue to next loop
	}
}
