package find_k_pairs_with_smallest_sums

import (
	"container/heap"
)

type MaxHeap [][]int

func (h MaxHeap) Less(i, j int) bool {
	return h[i][0]+h[i][1] > h[j][0]+h[j][1]
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

func RunSolution(nums1 []int, nums2 []int, k int) [][]int {
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)

	m := min(k, len(nums1))
	n := min(k, len(nums2))

	for i := range m {
		for j := range n {
			if maxHeap.Len() < k {
				heap.Push(&maxHeap, []int{nums1[i], nums2[j]})
				continue
			}

			if nums1[i]+nums2[j] < maxHeap.Peek()[0]+maxHeap.Peek()[1] {
				heap.Pop(&maxHeap)
				heap.Push(&maxHeap, []int{nums1[i], nums2[j]})
			}
		}
	}

	return maxHeap
}
