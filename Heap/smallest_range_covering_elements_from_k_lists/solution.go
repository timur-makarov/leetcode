package smallest_range_covering_elements_from_k_lists

import (
	"container/heap"
	"math"
)

type MinHeap [][]int

func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Peek() []int {
	return h[0]
}
func (h *MinHeap) Push(value interface{}) {
	*h = append(*h, value.([]int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[0 : n-1]
	return popped
}

func RunSolution(nums [][]int) []int {
	minHeap := MinHeap{}
	heap.Init(&minHeap)

	currentMaxVal := math.MinInt

	for i, row := range nums {
		heap.Push(&minHeap, []int{0, row[0], i})
		currentMaxVal = max(currentMaxVal, row[0])
	}

	wStart := 0
	wEnd := math.MaxInt

	for minHeap.Len() == len(nums) {
		item := heap.Pop(&minHeap).([]int)
		i, num, listI := item[0], item[1], item[2]

		if wEnd-wStart > currentMaxVal-num {
			wStart = num
			wEnd = currentMaxVal
		}

		if len(nums[listI])-1 > i {
			heap.Push(&minHeap, []int{i + 1, nums[listI][i+1], listI})
			currentMaxVal = max(currentMaxVal, nums[listI][i+1])
		}
	}

	return []int{wStart, wEnd}
}
