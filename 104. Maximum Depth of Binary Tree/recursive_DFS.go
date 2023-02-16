package _04__Maximum_Depth_of_Binary_Tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	return int(1+ math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}