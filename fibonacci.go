package main

import (
	"fmt"
)

/*
I will be coding both memory effiecient code & a short recrusive code for the fibonacci algorithm
*/

// Recursive Fib this function for a fibonacci algorithm is very short just returns with nth number
// however this is not memory optimized if you are working with limited memory
func fibRecur(num int) int {
	if num < 0 {
		// If the number input is anegative return -1 and print a msg
		fmt.Println("Error: can not be negative number")
		return -1
	}
	//is its 0 or 1 just return the number
	if num <= 1 {
		return num
	} else {
		// however is the number is 2 or greater we will recursively call the function
		return fibRecur(num-1) + fibRecur(num-2)
	}
}

/*
	fibIter is the memory optimized fib function little messy but this is what you use if memory is an issue
	this function just returns the nth number
*/
func fibIter(num int) int {
	//we create 2 var a & b to hold the first two numbers 0 & 1 fromthe fib sequence
	var a, b int = 0, 1

	//Temp to hold numbers as we traverse through the fib sequence
	var temp int

	//if negative num return -1 print msg
	if num < 0 {
		fmt.Println("Error: enter NonNegative Numbers")
		return -1
	}

	//is its 0 or 1 just return the number
	if num <= 1 {
		return num
	} else {

		for i := 2; i <= num; i++ {
			//we increment through the fib sequnce, we add the previous numbers together
			temp = a + b
			a = b
			b = temp //b will hold the current fib sew number
			fmt.Println(b)
		}
		return b
	}
}
