// Example: Switch

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ") // if, `i` = `1` then print 1 as `one`
	switch i {                     // basic switch:
	case 1: // case for switching
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { // switch: with functions
	case time.Saturday, time.Sunday: // you can use commas to use multi expressions in same case
		fmt.Println("It's, weekend!")
	default: // default is used for optional and rest return answer, if no one statement matches
		fmt.Println("It's, weekdays!")
	}

	t := time.Now()
	switch { // switch without an statement is like if/else statement
	case t.Hour() < 12: // case statement can be non-constant
		fmt.Println("It's before Noon!")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(false)
	whatAmI(5)
	whatAmI("hey")
}
