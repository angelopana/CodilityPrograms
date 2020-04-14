package main

import "fmt"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	stones = qsort(stones)
	for len(stones) > 0 {
		if stones[0] == stones[1] {
			stones = stones[2:] //destory max & 2nd max
		} else if stones[0] != stones[1] {
			stones[0] = stones[0] - stones[1]
			stones = append(stones[:1], stones[2:]...) //concat max and rest of number but the 2nd max
		}

		if len(stones) == 1 {
			return stones[0]
		}

		if len(stones) == 0 {
			return 0
		}
		stones = qsort(stones)
	}
	return -1
}

//quick sort thats ascending
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	// Pick a pivot
	// pivotIndex := rand.Int() % len(a)
	pivotIndex := (left + right) / 2
	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] > a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
	// sort first half and the second half
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func main() {
	test := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(test))
}
