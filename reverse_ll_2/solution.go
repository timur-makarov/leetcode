package reverse_ll_2

import (
	"github.com/timur-makarov/leetcode/linked_list"
)

/*
Given the head of a singly linked list and two integers left and right where left <= right,
reverse the nodes of the list from position left to position right, and return the reversed list.
*/

type ListNode = linked_list.ListNode

func RunSolution(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHead := &ListNode{}
	fakeHead.Next = head
	prev := fakeHead

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for i := left; i < right; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}

	return fakeHead.Next
}
