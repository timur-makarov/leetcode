package find_median_from_data_stream

import (
	"math"
)

type Heap struct {
	items    []int
	sortFunc func(a, b int) bool
}

func (h *Heap) Remove() *int {
	if len(h.items) == 0 {
		return nil
	}

	if len(h.items) == 1 {
		temp := h.items[0]
		h.items = []int{}
		return &temp
	}

	exMaxVal := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.SinkDown(0)

	return &exMaxVal
}

func (h *Heap) Insert(val int) {
	h.items = append(h.items, val)
	cur := len(h.items) - 1

	for cur > 0 && h.sortFunc(h.items[cur], h.items[h.ParentIndex(cur)]) {
		pi := h.ParentIndex(cur)
		h.Swap(pi, cur)
		cur = pi
	}
}

func (h *Heap) GetHeap() []int {
	return append([]int(nil), h.items...)
}

func (h *Heap) LeftChildIndex(i int) int {
	return 2*i + 1
}

func (h *Heap) RightChildIndex(i int) int {
	return 2*i + 2
}

func (h *Heap) ParentIndex(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

func (h *Heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) SinkDown(i int) {
	maxIndex := i
	size := len(h.items)

	for {
		li := h.LeftChildIndex(i)
		ri := h.RightChildIndex(i)

		if li < size && h.sortFunc(h.items[li], h.items[maxIndex]) {
			maxIndex = li
		}

		if ri < size && h.sortFunc(h.items[ri], h.items[maxIndex]) {
			maxIndex = ri
		}

		if maxIndex != i {
			h.Swap(i, maxIndex)
			i = maxIndex
		} else {
			break
		}
	}
}

type MedianFinder struct {
	maxHeap Heap
	minHeap Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: Heap{
			sortFunc: func(a, b int) bool {
				return a < b
			},
		},
		minHeap: Heap{
			sortFunc: func(a, b int) bool {
				return a > b
			},
		},
	}
}

func (f *MedianFinder) AddNum(num int) {
	if len(f.maxHeap.items) == 0 || f.maxHeap.items[0] <= num {
		f.maxHeap.Insert(num)
	} else {
		f.minHeap.Insert(num)
	}

	if len(f.maxHeap.items) > len(f.minHeap.items)+1 {
		f.minHeap.Insert(*f.maxHeap.Remove())
	} else if len(f.maxHeap.items) < len(f.minHeap.items) {
		f.maxHeap.Insert(*f.minHeap.Remove())
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.maxHeap.items) == len(f.minHeap.items) {
		return float64(f.maxHeap.items[0]+f.minHeap.items[0]) / 2
	}

	return float64(f.maxHeap.items[0])
}

func RunSolution(nums []int) float64 {
	obj := Constructor()

	for _, num := range nums {
		obj.AddNum(num)
	}

	return obj.FindMedian()
}
