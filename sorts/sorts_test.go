package sorts_test

import (
	"reflect"
	"testing"

	"github.com/Ahmed-Sermani/algos/sorts"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		seed   []int
		sorted []int
	}{
		{
			seed:   []int{5, 4, 3, 2, 1},
			sorted: []int{1, 2, 3, 4, 5},
		},
		{
			seed:   []int{32, 18, 22, 15, 35, 25, 21, 28, 28, 9, 3},
			sorted: []int{3, 9, 15, 18, 21, 22, 25, 28, 28, 32, 35},
		},
	}
	for _, test := range tests {
		s := sorts.HeapSort(test.seed)
		if !reflect.DeepEqual(test.sorted, s) {
			t.Logf("%+v != %+v", test.sorted, s)
			t.Fail()
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		seed   []int
		sorted []int
	}{
		{
			seed:   []int{5, 4, 3, 2, 1},
			sorted: []int{1, 2, 3, 4, 5},
		},
		{
			seed:   []int{32, 18, 22, 15, 35, 25, 21, 28, 28, 9, 3, 30},
			sorted: []int{3, 9, 15, 18, 21, 22, 25, 28, 28, 30, 32, 35},
		},
	}

	for _, test := range tests {
		s := sorts.QuickSort(test.seed)
		if !reflect.DeepEqual(test.sorted, s) {
			t.Logf("%+v != %+v", test.sorted, s)
			t.Fail()
		}
	}
}

func TestCountingSort(t *testing.T) {
	tests := []struct {
		seed   []int
		sorted []int
	}{
		{
			seed:   []int{5, 4, 3, 2, 1},
			sorted: []int{1, 2, 3, 4, 5},
		},
		{
			seed:   []int{32, 18, 22, 15, 35, 25, 21, 28, 28, 9, 3, 30},
			sorted: []int{3, 9, 15, 18, 21, 22, 25, 28, 28, 30, 32, 35},
		},
	}

	for _, test := range tests {
		s := sorts.CountingSort(test.seed)
		if !reflect.DeepEqual(test.sorted, s) {
			t.Logf("%+v != %+v", test.sorted, s)
			t.Fail()
		}
	}
}
