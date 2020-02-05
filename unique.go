package main

/*

Implement an alogirthm to determine if a string has all unique characters.

solution: By creating a hash map of hash[value]= true
		 if the hash map already has a char of true well return false, else keep adding into the map and set hash[value] to true

*/

// This function returns a bool if the stirng contains all unique characters in a string
func uniqueChar(s string) bool {
	//a map of bool to ensure rune == ascii so hashed char are insreted into the map and set to true
	hash := make(map[rune]bool)
	for _, char := range s {
		// if the hash[char] already exists in the map it means that it is not unique
		if _, ok := hash[char]; ok {
			return false
		}
		hash[char] = true
	}
	return true
}
