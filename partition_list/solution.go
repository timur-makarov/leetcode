package partition_list

import (
	"github.com/timur-makarov/leetcode/linked_list"
)

/*
Given the head of a linked list and a value x, partition it such that all nodes less than
x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
*/

type ListNode = linked_list.ListNode

func RunSolution(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHeadLess := &ListNode{}
	fakeHeadGreater := &ListNode{}

	less := fakeHeadLess
	greater := fakeHeadGreater

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}

		head = head.Next
	}

	greater.Next = nil
	less.Next = fakeHeadGreater.Next

	return fakeHeadLess.Next
}
