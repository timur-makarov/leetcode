package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/merge_k_sorted_lists"
	"slices"
	"testing"
)

func TestMergekSortedLists(t *testing.T) {
	head1 := &ListNode{}
	head2 := &ListNode{}
	head3 := &ListNode{}

	cur := head1
	for i := range 5 {
		cur.Next = &ListNode{Val: i * 1}
		cur = cur.Next
	}

	cur = head2
	for i := range 5 {
		cur.Next = &ListNode{Val: i * 2}
		cur = cur.Next
	}

	cur = head3
	for i := range 5 {
		cur.Next = &ListNode{Val: i * 3}
		cur = cur.Next
	}

	returnedHead := merge_k_sorted_lists.RunSolution([]*ListNode{head1.Next, head2.Next, head3.Next})
	var expectedSlice []int

	for i := range 5 {
		expectedSlice = append(expectedSlice, i*1)
	}
	for i := range 5 {
		expectedSlice = append(expectedSlice, i*2)
	}
	for i := range 5 {
		expectedSlice = append(expectedSlice, i*3)
	}

	slices.Sort(expectedSlice)

	assert.Equal(t, expectedSlice, listToSlice(returnedHead))
}
