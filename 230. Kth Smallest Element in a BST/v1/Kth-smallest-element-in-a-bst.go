/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	sorted := []int{}
	inorder(root, &sorted)
	return sorted[k-1]
}

func inorder(root *TreeNode, sorted *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, sorted)
	*sorted = append(*sorted, root.Val)
	inorder(root.Right, sorted)
}