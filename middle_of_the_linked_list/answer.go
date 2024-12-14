package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

*/

func Test() bool {
	head := &ListNode{-1, nil}
	cur := head

	for i := range 10 {
		cur.Next = &ListNode{i, nil}
		cur = cur.Next
	}

	midNode := middleNode(head)

	return midNode.Val == 4
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
