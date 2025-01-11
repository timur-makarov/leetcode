package kth_smallest_element_in_a_sorted_matrix

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Peek() int {
	return h[0]
}
func (h *MaxHeap) Push(value interface{}) {
	*h = append(*h, value.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[0 : n-1]
	return popped
}

func RunSolution(matrix [][]int, k int) int {
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)

	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			heap.Push(&maxHeap, matrix[i][j])

			if maxHeap.Len() > k {
				heap.Pop(&maxHeap)
			}
		}
	}

	return maxHeap.Peek()
}
