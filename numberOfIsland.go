package main

import (
	"fmt"
)

/*

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
	11110
	11010
	11000
	00000

Output: 1

Example 2:

Input:
	11000
	11000
	00100
	00011

Output: 3
*/

func numIslands(grid [][]byte) int {
	count := 0
	for row, i := range grid {
		for col, _ := range i {
			//we found a 1 as root now search around for other 1's to set to 0's
			if grid[row][col] == '1' {
				count++
				searchOtherOnes(grid, row, col) //search for other
			}
		}
	}
	return count
}

//searches for attachted '1's of the root island to set it to 0
func searchOtherOnes(grid [][]byte, i, j int) {
	//if we went over the matrix or found a 0 we are done return

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'
	//check bottom
	searchOtherOnes(grid, i+1, j)
	//check top
	searchOtherOnes(grid, i-1, j)
	//check left
	searchOtherOnes(grid, i, j-1)
	///check right
	searchOtherOnes(grid, i, j+1)
}

func main() {

	in := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(in))
}
