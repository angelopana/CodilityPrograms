/*

This is my lesson 1 codility answer, using Golang to solve the Binary Gap Problem


A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.
For example, number 9 has binary representation 1001 and contains a binary gap of length 2.
The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3.
The number 20 has binary representation 10100 and contains one binary gap of length 1.
The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.
For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.
Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:
		N is an integer within the range [1..2,147,483,647].


*/

package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	if N < 0 {
		return 0
	}
	num := int64(N)
	binaryNum := strconv.FormatInt(num, 2)
	zeroCounter := 0
	result := 0

	var check bool
	for _, i := range binaryNum {
		//start counting bool check
		if i == '1' {
			check = true
			if zeroCounter > result {
				result = zeroCounter
			}
			zeroCounter = 0
		} else if check {
			zeroCounter++
		}

	} //end forloop

	return result
}

func main() {

	/*

	   	For example, number 9 has binary representation 1001 and contains a binary gap of length 2.
	   The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3.
	   The number 20 has binary representation 10100 and contains one binary gap of length 1.
	   The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.
	*/

	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {

		input, _ := strconv.Atoi(reader.Text())

		fmt.Println(Solution(input))

	}

}
