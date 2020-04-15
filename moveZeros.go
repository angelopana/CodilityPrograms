package main

import "fmt"

/*

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:

    You must do this in-place without making a copy of the array.
    Minimize the total number of operations.
*/
func moveZeroes(nums []int) {

	for j, i := 0, 0; i < len(nums); i++ {
		fmt.Println(nums, i)
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		fmt.Println(nums, j)
	}
}

func main() {
	test := []int{0, 1, 0, 3, 12}
	moveZeroes(test)
}
