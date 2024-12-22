package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

Given the head of a linked list, remove the nth node from the end of the list and return its head.

*/

func Test() bool {
	head := &ListNode{-1, nil}
	//cur := head
	//
	//for i := range 1 {
	//	cur.Next = &ListNode{i, nil}
	//	cur = cur.Next
	//}

	head.Next = &ListNode{0, nil}

	returnedHead := runSolution(head, 3)

	return returnedHead.Val == 0
}

func runSolution(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	counter := 0

	for fast.Next != nil {
		fast = fast.Next

		if counter == n {
			slow = slow.Next
		} else {
			counter++
		}
	}

	if counter < n {
		return slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
