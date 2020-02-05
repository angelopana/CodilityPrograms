package main

import (
	"fmt"
	"strconv"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
	Input: 123
	Output: 321

Example 2:
	Input: -123
	Output: -321

Example 3:
	Input: 120
	Output: 21
*/

// Reverse function that given an integer it is reversed
func reverse(num string) {
	//converting a string into a num
	converted, err := strconv.Atoi(num)
	//must use err for Atoi
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Printf("inside funcition %d \n", converted)
}
