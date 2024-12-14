package main

import (
	"errors"
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length uint64
}

func main() {
	list := LinkedList{}

	list.Push(2)
	log.Println(list.String())

	list.Push(4)
	log.Println(list.String())

	log.Printf("Popped: %v\n", list.Pop())
	log.Println(list.String())

	list.Unshift(100)
	log.Println(list.String())

	log.Printf("Popped: %v\n", list.Pop())
	log.Println(list.String())

	list.Push(32)
	log.Println(list.String())

	log.Printf("Shifted: %v\n", list.Shift())
	log.Println(list.String())

	list.Push(32)
	list.Push(320)
	list.Push(3200)

	node, _ := list.Get(2)
	log.Printf("Got: %v\n", node)

	_ = list.Set(2, 322)
	log.Printf("Set node by the index of 2 to: %v\n", node)

	_ = list.Insert(0, 410)
	insertedNode, _ := list.Get(0)
	log.Printf("Inserted Node: %v\n", insertedNode)

	removedNode, _ := list.Remove(2)
	log.Printf("Removed Node: %v\n", removedNode)

	list.Reverse()
	log.Println(list.String())

	list.Reverse()
	log.Println(list.String())
}

func (list *LinkedList) String() string {
	return fmt.Sprintf("LinkedList: Head - %v  Tail - %v  Len - %v", list.head, list.tail, list.length)
}

func (list *LinkedList) Push(val int) {
	newNode := &ListNode{val, nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.Next = newNode
		list.tail = newNode
	}
	list.length++
}

func (list *LinkedList) Pop() *ListNode {
	if list.head == nil {
		return nil
	}

	temp := list.head
	prev := list.head

	for temp.Next != nil {
		prev = temp
		temp = temp.Next
	}

	list.tail = prev
	list.tail.Next = nil
	list.length--

	if list.length == 0 {
		list.head = nil
		list.tail = nil
	}

	return temp
}

func (list *LinkedList) Unshift(val int) {
	newNode := &ListNode{val, nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	}

	newNode.Next = list.head
	list.head = newNode
	list.length++
}

func (list *LinkedList) Shift() *ListNode {
	if list.head == nil {
		return nil
	}

	nodeToReturn := list.head

	if list.head.Next == nil {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.Next
	}

	list.length--
	return nodeToReturn
}

func (list *LinkedList) Get(i uint64) (*ListNode, error) {
	if list.head == nil {
		return nil, errors.New("list is empty")
	}

	if i >= list.length {
		return nil, errors.New("out of range")
	}

	cur := list.head

	for j := uint64(0); j < i; j++ {
		cur = cur.Next
	}

	return cur, nil
}

func (list *LinkedList) Set(i uint64, val int) error {
	if i == 0 {
		if list.head == nil {
			return errors.New("list is empty")
		}
		list.head.Val = val
		return nil
	}

	node, err := list.Get(i - 1)

	if err != nil {
		return err
	}

	node.Val = val

	return nil
}

func (list *LinkedList) Insert(i uint64, val int) error {
	if i == 0 {
		list.Unshift(val)
		return nil
	}

	if i == list.length {
		list.Push(val)
		return nil
	}

	newNode := &ListNode{val, nil}
	node, err := list.Get(i - 1)

	if err != nil {
		return err
	}

	newNode.Next = node.Next
	node.Next = newNode
	list.length++
	return nil
}

func (list *LinkedList) Remove(i uint64) (*ListNode, error) {
	if i == 0 {
		return list.Shift(), nil
	}

	if i == list.length {
		return list.Pop(), nil
	}

	node, err := list.Get(i - 1)

	if err != nil {
		return node, err
	}

	nodeToReturn := node.Next
	node.Next = node.Next.Next
	list.length--
	return nodeToReturn, nil
}

func (list *LinkedList) Reverse() {
	if list.head == nil {
		return
	}

	temp := list.head
	list.head = list.tail
	list.tail = temp
	next := temp.Next
	var prev *ListNode

	for i := uint64(0); i < list.length; i++ {
		next = temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
}
