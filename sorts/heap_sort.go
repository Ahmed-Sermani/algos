package sorts

import "github.com/Ahmed-Sermani/algos/structures"

func HeapSort(arr []int) []int {
	sorted := make([]int, len(arr))
	h := structures.NewHeapFromTree(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		sorted[i] = h.PopRoot()
	}
	return sorted
}
