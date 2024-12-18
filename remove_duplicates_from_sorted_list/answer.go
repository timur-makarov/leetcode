package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.

*/

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

func RunSolution2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHead := &ListNode{}
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
