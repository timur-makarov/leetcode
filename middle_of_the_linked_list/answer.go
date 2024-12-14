package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

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
