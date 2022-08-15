package Lowest_Common_Ancestor_of_a_BST
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for{
		switch{
		case root.Val == p.Val || root.Val == q.Val:    //LCA is p or q
			return root
		case root.Val > p.Val && root.Val > q.Val:      //p,q on the left
			root = root.Left
		case root.Val < p.Val && root.Val < q.Val:      //p,q on the right
			root = root.Right
		default:
			return root                                 //p,q on different side
		}
	}
	return root
}