# 230. Kth Smallest Element in a BST

## v1

v1 use inorder traversal to create a sorted slice(array), return the Kth element.

use additional memory

Time complexity: O(n)  
Memory complexity: O(n)  
n is the size of the tree

## v2

v2 use a counter to stop at the Kth node, and return the value.

use 2 int variable.

This is faster than v1 and use less memory.

Time complexity: O(k)
