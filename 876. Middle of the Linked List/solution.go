package _76__Middle_of_the_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "math"

func middleNode(head *ListNode) *ListNode {
	var n = 1.0
	ans := head
	temp := head
	for temp.Next != nil {
		n++
		temp = temp.Next
	}
	if(math.Ceil(n/2)!=(n/2)){
		for i:= 0; i< int( math.Ceil(n / 2) -1);i++{

			ans = ans.Next
		}

	}else{
		for i:= 0; i< int(math.Ceil(n / 2));i++{

			ans = ans.Next
		}


	}

	return ans

}