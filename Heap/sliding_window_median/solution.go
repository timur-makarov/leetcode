package sliding_window_median

import "container/heap"

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Peek() int {
	return (*h)[0]
}
func (h *MinHeap) Push(value interface{}) {
	*h = append(*h, value.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[0 : n-1]
	return popped
}

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *MaxHeap) Len() int {
	return len(*h)
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Peek() int {
	return (*h)[0]
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

func RunSolution(nums []int, k int) []float64 {
	var result []float64

	maxHeap, minHeap := &MaxHeap{}, &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	toDelete := make(map[int]int, 0)

	for i := 0; i < k; i++ {
		heap.Push(maxHeap, nums[i])
	}

	for i := 0; i < k/2; i++ {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}

	isOdd := k%2 != 0

	for i := k; i <= len(nums); i++ {
		heapsBalance := 0

		if isOdd {
			result = append(result, float64(maxHeap.Peek()))
		} else {
			result = append(result, float64(maxHeap.Peek()+minHeap.Peek())/2)
		}

		if i == len(nums) {
			break
		}

		in := nums[i]
		out := nums[i-k]

		if maxHeap.Peek() >= out {
			heapsBalance--
		} else {
			heapsBalance++
		}

		toDelete[out]++

		if maxHeap.Peek() >= in {
			heap.Push(maxHeap, in)
			heapsBalance++
		} else {
			heap.Push(minHeap, in)
			heapsBalance--
		}

		if heapsBalance > 0 {
			heap.Push(minHeap, heap.Pop(maxHeap))
		} else if heapsBalance < 0 {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}

		for toDelete[maxHeap.Peek()] > 0 {
			toDelete[heap.Pop(maxHeap).(int)]--
		}
		for minHeap.Len() > 0 && toDelete[minHeap.Peek()] > 0 {
			toDelete[heap.Pop(minHeap).(int)]--
		}
	}

	return result
}
