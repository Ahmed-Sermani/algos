package structures

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue([]int{20, 10, 10, 4, 5, 6, 3, 1})
	if m := pq.Max(); m != 20 {
		t.Logf("%d != %d", m, 20)
		t.Fail()
	}
	if m := pq.ExtractMax(); m != 20 {
		t.Fail()
	}
	if !reflect.DeepEqual(pq.h.store[:7], []int{10, 5, 10, 4, 1, 6, 3}) {
		t.Fail()
	}
	if pq.h.size != 7 {
		t.Fail()
	}
	f := func() {
		defer func() {
			if r := recover(); r == nil {
				t.Log("panic expected")
				t.Fail()
			}
		}()
		pq.IncreaseKey(0, 1)
	}
	f()
	pq.IncreaseKey(1, 11)
	if !reflect.DeepEqual(pq.h.store[:7], []int{11, 10, 10, 4, 1, 6, 3}) {
		t.Fail()
	}
	pq.Insert(5)
	if pq.h.store[3] != 5 {
		t.Fail()
	}
}

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Panics(t, func() { stack.Pop() })
}

func TestQueue(t *testing.T) {
	queue := NewQueue[string](4)
	queue.Enqueue("1")
	queue.Enqueue("2")
	assert.Equal(t, "1", queue.Dequeue())
	queue.Enqueue("3")
	queue.Enqueue("4")
	queue.Enqueue("5")
	assert.Panics(t, func() { queue.Enqueue("6") })
	assert.Equal(t, "2", queue.Dequeue())
	assert.Equal(t, "3", queue.Dequeue())
	assert.Equal(t, "4", queue.Dequeue())
	assert.Equal(t, "5", queue.Dequeue())
	assert.Panics(t, func() { queue.Dequeue() })
}
