// Write a test for the linked list

package main

import (
	"testing"

	"github.com/timur-makarov/leetcode/linked_list"
)

type ListNode = linked_list.ListNode

func TestLinkedList(t *testing.T) {
	linked_list.RunLinkedList()
}
