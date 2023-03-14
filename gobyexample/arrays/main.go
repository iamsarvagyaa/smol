package main

import "fmt"

func main() {
	var a [5]int            // here we are declaring an array of 5 elements (but it's a one dimensional)
	fmt.Println("emp: ", a) // here we are printing the array

	a[4] = 100 // we can assign an array via its position
	fmt.Println(a)
	fmt.Println(a[4])   // can also print the element via position of an array
	fmt.Println(len(a)) // len is to check the length of array

	b := [5]int{1, 2, 3, 4, 5} // here we are declaring the arrays
	fmt.Println(b)

	var twoD [2][3]int // 2 rows and 3 columns, we declared an empty array
	fmt.Println(twoD)

	for i := 0; i < 2; i++ { // the first loop
		fmt.Println(i)
		// fmt.Println(twoD)

		for j := 0; j < 3; j++ { // the second nested loop
			fmt.Println(j)
			// fmt.Println(twoD)
			twoD[i][j] = i + j // so, this will add in loop array

			// 0 (from first loop)
			// 0 (from nested loop)
			// add: 0+0 = 0,
			// 0 (the first loop again because, right now we are in second loop, the nested loop have, true value)
			// 1 (from nested loop)
			// add: 0+1 = 1,
			// 0 (again the first loop value)
			// 2 (from nested loop)
			// add: 0+2 = 2,
			// now the nested loop, statement got completed, can move to first loop now
			// 1 (from first loop)
			// 0 (from nested loop)
			// add: 1+0 = 1,
			// 1 (from first loop)
			// 1 (from nested loop
			// add: 1+1 = 2,
			// 1 (from first loop)
			// 2 (from nested loop, now again the nested loop statement completed. completed both loop. it will stop now)
			// add: 1+2 = 3

			// [[0, 1, 2] [1, 2, 3]]
		}
	}

	fmt.Println(twoD)
}
