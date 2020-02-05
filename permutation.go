package main

/*
Purpose: Given two strings, write a method to decide if one is a permutation of the other

solution: We create a hash map of value int, if we inserte hash[charValue] and increment it per each char seen
		  we then loop through the second B string, into the hash map and decrement `--`, and it anyhting is
		  less than 0 it means that char was never in the permuation string, and that value is negative
		  else everything is 0 and returns true
*/
func permutation(s, a string) bool {
	// check length if its correct if not false
	if len(a) != len(s) {
		return false
	}
	// create a hash map
	hash := make(map[rune]int)
	for _, charS := range s {
		//input the char value into the string hash[map] and increment to show there is either 1+ characters in the string
		hash[charS]++
	}

	// This little forloop subtracts from the either permuation stirng or stirng B
	// if any of the values are negative it means that the char value was never there and returns false
	for _, charA := range a {
		hash[charA]--
		if hash[charA] < 0 {
			return false
		}
	}
	return true //else everything is good and return true
}
