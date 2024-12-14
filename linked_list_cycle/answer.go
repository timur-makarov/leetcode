package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached
again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.

*/

func Test() bool {
	head := &ListNode{-1, nil}
	cur := head

	for i := range 5 {
		cur.Next = &ListNode{i, nil}
		cur = cur.Next
	}

	//cur.Next = head

	return hasCycle(head) == false
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
