/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return pathSumHelper(root, targetSum, 0)
}

func pathSumHelper(root *TreeNode, targetSum, currentSum int) bool {
	if root == nil {
		return false
	}

	currentSum += root.Val

	if root.Left == nil && root.Right == nil && currentSum == targetSum {
		return true
	}
	if pathSumHelper(root.Left, targetSum, currentSum) {
		return true
	}
	if pathSumHelper(root.Right, targetSum, currentSum) {
		return true
	}

	return false
}