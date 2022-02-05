package structures

type Heap struct {
	size  int
	store []int
}

func NewHeapFromTree(source []int) *Heap {
	h := &Heap{
		size:  len(source),
		store: source,
	}
	// elements from source[size/2:n] are leaves.
	// so we only run the heapify subroutine as for the first half.
	for i := (h.size / 2) - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h *Heap) MaxHeapify(i int) {
	largest := i
	for {
		l := left(i)
		r := right(i)
		if l < h.size && h.store[l] > h.store[i] {
			largest = l
		}
		if r < h.size && h.store[r] > h.store[largest] {
			largest = r
		}
		if largest == i {
			return
		}
		h.swap(i, largest)
		i = largest
	}
}

func (h *Heap) UpHeapify(i int) {
	for i > 0 && h.store[i] > h.store[parent(i)] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *Heap) PopRoot() int {
	if h.IsEmpty() {
		panic("heap is empty")
	}
	root := h.store[0]
	h.store[0] = h.store[h.size-1]
	h.size--
	h.MaxHeapify(0)
	return root
}

func (h *Heap) IsEmpty() bool {
	return len(h.store) == 0
}

func (h *Heap) swap(i, j int) {
	h.store[i], h.store[j] = h.store[j], h.store[i]
}

func left(i int) int {
	return (i << 1) + 1
}

func right(i int) int {
	return (i << 1) + 2
}

func parent(i int) int {
	return (i - 1) >> 1
}
