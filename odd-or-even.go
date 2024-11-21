package main

import (
	"fmt"
)

// OddOrEven takes a number and returns whether it is odd or even
// Create a slice of ints from 0 through 10
// Iterate through the slice
// If the number is even, print "even"

func OddOrEven() {
	nums := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, num := range nums {
		if num % 2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}