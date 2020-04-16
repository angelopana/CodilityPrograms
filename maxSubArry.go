package main

import "fmt"
/* 

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


*/

func maxSubArray(nums []int) int {
    
    if len(nums) == 1 {
        return nums[0]
    }
    
    sums := nums[0]
    max :=  nums[0]
    
    
    for i := 1; i < len(nums); i++{
        if (sums + nums[i]) > nums[i] {
            sums += nums[i]
        } else {
            sums = nums[i]
        }
        
        if sums > max {
            max = sums
        }
    }
    
    return max
    
func main() {

	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	test := []int{-2, 1}

	fmt.Println(maxSubArray(test))
}
