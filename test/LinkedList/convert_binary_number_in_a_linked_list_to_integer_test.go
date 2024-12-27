package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/LinkedList/convert_binary_number_in_a_linked_list_to_integer"
	"testing"
)

func TestConvertBinaryNumberInALinkedListToInteger(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 0}
	head.Next.Next = &ListNode{Val: 1}

	res := convert_binary_number_in_a_linked_list_to_integer.RunSolution(head)

	assert.Equal(t, 5, res)
}
