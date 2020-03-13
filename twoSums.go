package main

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].


*/

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	var sums []int
	for i, numbers := range nums {
		dif := target - numbers
		if _, ok := hash[dif]; ok {
			sums = append(sums, i, hash[dif])
		}
		//hash[value] = index
		hash[numbers] = i
	}

	return sums
}
