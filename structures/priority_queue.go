package structures

import "math"

type PriorityQueue struct {
	h *Heap
}

func NewPriorityQueue(source []int) *PriorityQueue {
	return &PriorityQueue{
		h: NewHeapFromTree(source),
	}
}
func (q *PriorityQueue) Max() int {
	if q.h.IsEmpty() {
		panic("p-queue is empty")
	}
	return q.h.store[0]
}

func (q *PriorityQueue) ExtractMax() int {
	if q.h.IsEmpty() {
		panic("underflow")
	}

	return q.h.PopRoot()
}

func (q *PriorityQueue) IncreaseKey(i, k int) {
	if q.h.store[i] > k {
		panic("key less than the current")
	}
	q.h.store[i] = k
	q.h.UpHeapify(i)
}

func (q *PriorityQueue) Insert(key int) {
	q.h.size++
	if len(q.h.store) < q.h.size {
		q.h.store = append(q.h.store, math.MinInt)
	} else {
		q.h.store[q.h.size-1] = math.MinInt
	}
	q.IncreaseKey(q.h.size-1, key)
}
