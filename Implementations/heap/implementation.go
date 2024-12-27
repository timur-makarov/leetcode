package heap

import (
	"log"
	"math"
)

type Heap struct {
	items []int
}

func RunHeapFunctions() {
	var heap Heap

	heap.Insert(1)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(5)

	log.Println(heap.items)

	heap.Insert(12)
	heap.Insert(15)
	log.Println(heap.items)

	heap.Remove()
	heap.Remove()
	log.Println(heap.items)
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

	for cur > 0 && h.items[cur] > h.items[h.ParentIndex(cur)] {
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

		if li < size && h.items[li] > h.items[maxIndex] {
			maxIndex = li
		}

		if ri < size && h.items[ri] > h.items[maxIndex] {
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
