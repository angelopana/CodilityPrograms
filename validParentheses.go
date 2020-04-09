package main

/*

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.


Example 1:
	Input: "()"
	Output: true

Example 2:
	Input: "()[]{}"
	Output: true

Example 3:
	Input: "(]"
	Output: false

Example 4:
	Input: "([)]"
	Output: false

Example 5:
	Input: "{[]}"
	Output: true
*/

func isValid(s string) bool {
	var stack []rune
	pMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, p := range s {
		if p == '(' || p == '[' || p == '{' {
			// add to stack
			stack = append(stack, p) //i.e the push onto stack
		}
		if _, ok := pMap[p]; ok {
			if stack[len(stack)-1] == pMap[p] {
				stack = stack[:len(stack)-1] //pops from the stack LIFO
			} else {
				return false
			}
		}
	}

	// if stack is empty all parentheses are valid
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
