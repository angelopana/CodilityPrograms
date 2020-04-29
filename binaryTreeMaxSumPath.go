package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//this is for negative numbers
	ans := root.Val
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left_Sum := max(helper(root.Left, sum), 0)
	right_Sum := max(helper(root.Right, sum), 0)

	temp := left_Sum + right_Sum + root.Val
	*sum = max(temp, *sum)

	return max(left_Sum, right_Sum) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
