package main

import "fmt"

/*

Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

Example 1:

Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.

Example 2:

Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Note: The length of the given binary array will not exceed 50,000. */

func findMaxLength(nums []int) int {
	hash := make(map[int]int)
	var temp, count, max int
	hash[0] = -1
	for index, number := range nums {
		if number == 0 {
			count--
		}
		if number == 1 {
			count++
		}
		//the count is already in the hash
		if hashVal, ok := hash[count]; ok {
			temp = index - hashVal
			if temp > max {
				max = temp
			}
		} else {
			hash[count] = index
		}
	}
	return max
}

func main() {

	test := []int{0, 1}
	fmt.Println(findMaxLength(test))
}
