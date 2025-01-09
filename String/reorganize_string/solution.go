package reorganize_string

import (
	"container/heap"
	"strings"
)

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

func RunSolution(s string) string {
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)

	var result strings.Builder

	freqs := make(map[int]int, 26)
	for _, l := range s {
		freqs[int(l)]++
	}

	if len(freqs) == 1 {
		return ""
	}

	for l, freq := range freqs {
		heap.Push(&maxHeap, []int{l, freq})
	}

	prevL := -1
	prevFreq := 0

	for maxHeap.Len() != 0 {
		item, _ := heap.Pop(&maxHeap).([]int)
		l, freq := item[0], item[1]

		if prevL != -1 && prevFreq > 0 {
			heap.Push(&maxHeap, []int{prevL, prevFreq})
		}

		result.WriteByte(byte(l))
		prevL = l
		prevFreq = freq - 1
	}

	if result.Len() != len(s) {
		return ""
	}

	return result.String()
}
