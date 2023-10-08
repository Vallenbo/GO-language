package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	num := []int{}
	if root == nil {
		return 0
	}
	num = append(num, root.Val)

	if root.Left != nil {
		num = append(num, root.Left.Val)
	} else {
		num = append(num)
	}

	if root.Right != nil {
		num = append(num, root.Right.Val)

	}
	return 0
}

func main() {

}
