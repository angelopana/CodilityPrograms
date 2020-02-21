package main

/*
Given a non-negative integer num, return the number of steps to reduce it to zero.
If the current number is even, you have to divide it by 2,
otherwise, you have to subtract 1 from it.
*/

func check(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if check(num) {
			num = num / 2
		} else {
			num = num - 1
		}
		count++
	}
	return count
}
