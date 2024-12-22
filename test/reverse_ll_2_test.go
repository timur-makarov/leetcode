package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/reverse_ll_2"
	"testing"
)

func TestReverseLL2(t *testing.T) {
	head := &ListNode{}

	cur := head
	for i := range 10 {
		cur.Next = &ListNode{Val: i + 1}
		cur = cur.Next
	}

	t.Log(listToSlice(head))

	left := 3
	right := 7
	returnedHead := reverse_ll_2.RunSolution(head, left, right)

	expectedSlice := []int{0, 1, 2, 7, 6, 5, 4, 3, 8, 9, 10}

	t.Log(listToSlice(returnedHead))
	assert.Equal(t, expectedSlice, listToSlice(returnedHead))
}
