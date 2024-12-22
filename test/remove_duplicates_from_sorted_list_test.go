package main

import (
	"github.com/timur-makarov/leetcode/linked_list"
	"github.com/timur-makarov/leetcode/remove_duplicates_from_sorted_list"
	"testing"
)

type ListNode = linked_list.ListNode

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	head := &ListNode{}

	cur := head
	for i := range 10 {
		cur.Next = &ListNode{Val: i + 1}
		cur = cur.Next
	}

	i := 0
	cur = head
	for cur.Next != nil {
		if i != 0 && i%3 == 0 {
			cur.Val = i
			cur.Next.Val = i
		}

		cur = cur.Next
		i++
	}

	var listAsSlice []int

	cur = head
	for cur.Next != nil {
		listAsSlice = append(listAsSlice, cur.Val)
		cur = cur.Next
	}

	t.Log(listAsSlice) // 0 1 2 3 3 5 6 6 8 9

	returnedHead := remove_duplicates_from_sorted_list.RunSolution(head)

	sieve := make(map[int]bool)

	cur = returnedHead

	for cur.Next != nil {
		if ok := sieve[cur.Val]; ok {
			t.Fatal("duplicate element found")
			return
		}

		cur = cur.Next
	}
}

func TestRemoveDuplicatesFromSortedList2(t *testing.T) {
	head := &ListNode{}

	cur := head
	for i := range 13 {
		cur.Next = &ListNode{Val: i + 1}
		cur = cur.Next
	}

	i := 0
	cur = head
	for cur.Next.Next != nil {
		if i != 0 && i%4 == 0 {
			cur.Val = i
			cur.Next.Val = i
			cur.Next.Next.Val = i
		}

		cur = cur.Next
		i++
	}

	var listAsSlice []int

	cur = head
	for cur.Next != nil {
		listAsSlice = append(listAsSlice, cur.Val)
		cur = cur.Next
	}

	t.Log(listAsSlice) // 0 1 2 3 4 4 4 7 8 8 8 11 12

	returnedHead := remove_duplicates_from_sorted_list.RunSolution2(head)

	sieve := make(map[int]bool)
	sieve[4] = true
	sieve[8] = true

	cur = returnedHead

	for cur.Next != nil {
		if ok := sieve[cur.Val]; ok {
			t.Fatal("failed the task")
			return
		}

		cur = cur.Next
	}
}
