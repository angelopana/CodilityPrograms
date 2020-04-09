package main

import (
	"fmt"
	"sort"
)

/*

Problem:
	Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"]
Output:
	[
		["ate","eat","tea"],
		["nat","tan"],
		["bat"]
	]

Note:

    All inputs will be in lowercase.
    The order of your output does not matter.


*/
func groupAnagrams(strs []string) [][]string {

	hash := make(map[string][]string)
	for _, value := range strs {
		// temp :=
		// if _, ok := hash[temp]; ok {
		// 	hash[temp] = append(hash[temp], value)
		// }
		hash[sortRunes(value)] = append(hash[sortRunes(value)], value)

	}

	var output [][]string
	for _, i := range hash {
		output = append(output, i)
		fmt.Println(output)
	}

	return output
}

// This sorts the char/runs
func sortRunes(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {

	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	i := groupAnagrams(input)

	for _, n := range i {
		fmt.Println(n)
	}
}
