package structures

import (
	"reflect"
	"testing"
)

func TestMaxHeapify(t *testing.T) {
	t.Parallel()
	tests := []struct {
		seed     []int
		idx      int
		expected []int
	}{
		{
			seed:     []int{1, 2},
			idx:      0,
			expected: []int{2, 1},
		},
		{
			seed:     []int{10, 20, 3, 4, 5, 6, 10, 1},
			idx:      0,
			expected: []int{20, 10, 10, 4, 5, 6, 3, 1},
		},
	}
	for _, test := range tests {
		h := NewHeapFromTree(test.seed)
		if !reflect.DeepEqual(test.expected, h.store) {
			t.Logf("%v != %v", h.store, test.expected)
			t.Fail()
		}
	}
}

func TestHeapPopRoot(t *testing.T) {
	t.Parallel()
	tests := []struct {
		heap     []int
		root     int
		afterPop []int
	}{
		{
			heap:     []int{2, 1},
			root:     2,
			afterPop: []int{1},
		},
		{
			heap:     []int{20, 10, 10, 4, 5, 6, 3, 1},
			root:     20,
			afterPop: []int{10, 10, 4, 6, 6, 3, 1},
		},
		{
			heap:     []int{1},
			root:     1,
			afterPop: []int{},
		},
	}

	for _, test := range tests {
		h := &Heap{store: test.heap, size: len(test.heap)}
		r := h.PopRoot()
		if r != test.root {
			t.Logf("%d != %d", r, test.root)
			t.Fail()
		}
		if !reflect.DeepEqual(h.store, test.afterPop) {
			t.Logf("%+v != %+v", h.store, test.afterPop)
		}
	}
}
