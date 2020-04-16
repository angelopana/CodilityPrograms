package main

import "fmt"

func countElements(arr []int) int {
	hash := make(map[int]bool)
	for _, n := range arr {
		hash[n] = true
	}

	count := 0
	for _, i := range arr {
		//in the hash table count
		if _, ok := hash[i+1]; ok {
			count++
		}
	}
	return count
}

func main() {

	test := []int{1, 3, 2, 3, 5, 0}
	// t := []int{1, 2, 3}
	fmt.Println("\n", countElements(test))
}
