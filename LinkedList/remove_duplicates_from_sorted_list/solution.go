package remove_duplicates_from_sorted_list

import (
	"github.com/timur-makarov/leetcode/Implementations/linked_list"
)

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
*/

type ListNode = linked_list.ListNode

func RunSolution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}

		cur = cur.Next
	}

	return head
}
