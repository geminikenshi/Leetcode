package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}

	xi := []int{}
	xi = traversal(root, xi)

	return xi
}

func traversal(root *TreeNode, xi []int) []int{
	if root == nil {
		return nil
	}
	if root.Left != nil{
		xi = traversal(root.Left, xi)
	}
	xi = append(xi, root.Val)
	if root.Right != nil{
		xi = traversal(root.Right, xi)
	}

	return xi
}