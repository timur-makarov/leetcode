package merge_k_sorted_lists

import (
	"container/heap"
	"github.com/timur-makarov/leetcode/Implementations/linked_list"
)

type MinHeap []*ListNode

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Peek() *ListNode {
	return h[0]
}
func (h *MinHeap) Push(value interface{}) {
	*h = append(*h, value.(*ListNode))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[0 : n-1]
	return popped
}

type ListNode = linked_list.ListNode

func RunSolution(lists []*ListNode) *ListNode {
	resultHead := &ListNode{}

	minHeap := MinHeap{}
	heap.Init(&minHeap)

	for _, list := range lists {
		heap.Push(&minHeap, list)
	}

	cur := resultHead

	for minHeap.Len() != 0 {
		node := heap.Pop(&minHeap).(*ListNode)

		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
	}

	return resultHead.Next
}
