package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//This is from LeetCode 938 the struct given is above from leet code
//will not work for the actual main package on go
func rangeSumBST(root *TreeNode, L int, R int) int {
	// var counter int
	counter := 0
	if root == nil {
		return 0
	}

	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}

	counter += root.Val
	counter += rangeSumBST(root.Left, L, R)
	counter += rangeSumBST(root.Right, L, R)
	return counter

}
