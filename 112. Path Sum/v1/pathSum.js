/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} targetSum
 * @return {boolean}
 */
var hasPathSum = function (root, targetSum) {
  if (!root) {
    return false;
  }
  return pathSumHelper(root, targetSum, 0);
};

function pathSumHelper(root, targetSum, currentSum) {
  if (!root) {
    return false;
  }
  currentSum = currentSum + root.val;
  if (!root.left && !root.right) {
    return targetSum === currentSum;
  }
  if (pathSumHelper(root.left, targetSum, currentSum)) {
    return true;
  }
  if (pathSumHelper(root.right, targetSum, currentSum)) {
    return true;
  }

  return false;
}
