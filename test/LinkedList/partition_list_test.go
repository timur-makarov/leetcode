package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/LinkedList/partition_list"
	"testing"
)

func TestPartitionList(t *testing.T) {
	head := &ListNode{}

	cur := head
	for i := range 10 {
		cur.Next = &ListNode{Val: i + 1}
		cur = cur.Next
	}

	i := 0
	cur = head
	for cur != nil {
		if i != 0 && i%3 != 0 {
			cur.Val = 10
		}

		cur = cur.Next
		i++
	}

	t.Log(listToSlice(head))

	x := 6
	returnedHead := partition_list.RunSolution(head, x)

	expectedSlice := []int{0, 3, 10, 10, 10, 10, 6, 10, 10, 9, 10}

	t.Log(listToSlice(returnedHead))
	assert.Equal(t, expectedSlice, listToSlice(returnedHead))
}

func listToSlice(head *ListNode) []int {
	var listAsSlice []int

	cur := head
	for cur != nil {
		listAsSlice = append(listAsSlice, cur.Val)
		cur = cur.Next
	}

	return listAsSlice
}
