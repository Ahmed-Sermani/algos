package sorts

import "math"

// counting sort is a linear sorting algorithms that
// doesn't use comparsions. instead it exploits some information in
// the data set.
// here give that the max number in A is k. k is no larger than n^2 then
// sorting in linear time is possable.
// the space complexity is O(k + n)
func CountingSort(A []int) []int {
	idx := make([]int, getMax(A)+1)
	for i := range A {
		idx[A[i]]++
	}
	for i := 1; i < len(idx); i++ {
		idx[i] += idx[i-1]
	}
	sorted := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		idx[A[i]]--
		pos := idx[A[i]]
		sorted[pos] = A[i]
	}
	return sorted
}

func getMax(A []int) int {
	max := math.MinInt
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}
