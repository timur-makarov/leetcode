package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

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
