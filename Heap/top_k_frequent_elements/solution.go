package top_k_frequent_elements

import "container/heap"

type MaxHeap [][]int

func (h MaxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}
func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Peek() []int {
	return h[0]
}
func (h *MaxHeap) Push(value interface{}) {
	*h = append(*h, value.([]int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[0 : n-1]
	return popped
}

func RunSolution(nums []int, k int) []int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	hm := make(map[int]int)

	for _, num := range nums {
		hm[num]++
	}

	for num, freq := range hm {
		heap.Push(maxHeap, []int{num, freq})
	}

	var result []int

	for len(result) != k {
		result = append(result, heap.Pop(maxHeap).([]int)[0])
	}

	return result
}
