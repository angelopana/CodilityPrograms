package main

import "fmt"

/*
	Given an array nums of n integers where n > 1,
	return an array output such that output[i]
	is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]

Constraint:
	It's guaranteed that the product of the elements of any prefix or suffix
	of the array (including the whole array) fits in a 32 bit integer.

	Note: Please solve it without division and in O(n).

Follow up:
	Could you solve it with constant space complexity?
	(The output array does not count as extra space for the purpose of space complexity analysis.)


*/

func productExceptSelf(nums []int) []int {

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	//tail for front
	left[0] = 1
	//tail in tha bck
	right[len(nums)-1] = 1

	//idea is to loop from front to back
	// test := [1,2,3,4] > [1,1,2,6]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	//loop from back to front
	// test := [1,2,3,4] > [24,12,4,1]
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}
	fmt.Println(left, right)
	output := make([]int, len(nums))
	// As we created the left(front to back) & right(back to front) arrays
	// we multiple the given elements to left[i] * right[i]
	/*
		i.e
			[1,1,2,6] * [24,12,4,1] => [24,12,8,6]

	*/
	for i := 0; i < len(nums); i++ {
		output[i] = left[i] * right[i]
	}

	return output
}

func main() {

	test := []int{2, 5, 3, 4, 1}

	fmt.Println(productExceptSelf(test))

}
