package remove_duplicates_from_sorted_list

import (
	"github.com/timur-makarov/leetcode/Implementations/linked_list"
)

/*

Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
Return the linked list sorted as well.

*/

func RunSolution2(head *linked_list.ListNode) *linked_list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHead := &linked_list.ListNode{}
	fakeHead.Next = head
	cur := fakeHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			sameVal := cur.Next.Val

			for cur.Next != nil && cur.Next.Val == sameVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return fakeHead.Next
}
