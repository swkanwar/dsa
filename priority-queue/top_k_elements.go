package main

import (
	"container/heap"
	"fmt"
)

type entry struct {
	value    int
	priority int
}

// EntryHeap is a min-heap of entries.
type EntryHeap []entry

func (h EntryHeap) Len() int           { return len(h) }
func (h EntryHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h EntryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EntryHeap) Push(x interface{}) {
	*h = append(*h, x.(entry))
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TopK returns the k largest elements from the input slice.
func TopK(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	h := &EntryHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, entry{value: num, priority: num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	n := h.Len()
	result := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(entry).value
	}
	return result
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Printf("Top %d elements: %v\n", k, TopK(nums, k))
}
