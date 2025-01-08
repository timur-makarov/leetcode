package k_closest_points_to_origin

import (
	"container/heap"
)

type MaxHeap [][]int

func (h *MaxHeap) Less(i, j int) bool {
	return distanceFromOrigin((*h)[i]) > distanceFromOrigin((*h)[j])
}
func (h *MaxHeap) Len() int {
	return len(*h)
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Peek() []int {
	return (*h)[0]
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

func distanceFromOrigin(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func RunSolution(points [][]int, k int) [][]int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for i := range k {
		heap.Push(maxHeap, points[i])
	}

	for i := k; i < len(points); i++ {
		if distanceFromOrigin(points[i]) < distanceFromOrigin(maxHeap.Peek()) {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, points[i])
		}
	}

	return *maxHeap
}
