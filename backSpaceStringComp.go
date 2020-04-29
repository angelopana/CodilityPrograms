package main

import "fmt"

/*

Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".

Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".

Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".


*/

func backspaceCompare(S string, T string) bool {

	var s1 []rune
	var s2 []rune

	s1 = addToStack(S)
	s2 = addToStack(T)

	if string(s1) == string(s2) {
		return true
	}
	return false
}

func addToStack(s string) []rune {
	var stack []rune
	for _, j := range s {
		if j == '#' {
			if len(stack) != 0 {
				top := len(stack) - 1
				stack = stack[:top] // pop
			}

		} else {
			stack = append(stack, j)
		}
	}
	return stack
}

func main() {

	s := "a##c"
	t := "#a#c"

	fmt.Println(backspaceCompare(s, t))
}
