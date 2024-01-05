/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	counter := 0
	ans := 0
	inorder(root, k, &counter, &ans)
	return ans
}

func inorder(root *TreeNode, k int, counter, ans *int) {
	if root == nil {
		return
	}
	inorder(root.Left, k, counter, ans)

	if *counter++; *counter == k {
		*ans = root.Val
		return
	}
	inorder(root.Right, k, counter, ans)
}