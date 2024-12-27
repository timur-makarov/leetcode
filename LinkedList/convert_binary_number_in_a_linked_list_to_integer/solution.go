package convert_binary_number_in_a_linked_list_to_integer

import (
	"github.com/timur-makarov/leetcode/Implementations/linked_list"
)

/*

Given head which is a reference node to a singly-linked list.
The value of each node in the linked list is either 0 or 1.
The linked list holds the binary representation of a number.
Return the decimal value of the number in the linked list.
The most significant bit is at the head of the linked list.

*/

func RunSolution(head *linked_list.ListNode) int {
	if head == nil {
		return 0
	}

	if head.Next == nil {
		return head.Val
	}

	res := 0
	cur := head

	for cur != nil {
		res = res*2 + cur.Val
		cur = cur.Next
	}

	return res
}
