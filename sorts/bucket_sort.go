package sorts

import "math"

// bucket sort is non-comparsion based sort. As the rest of its kind it expolits
// some information about the data. here it assumes that the input is unifromally distributed
func BucketSort(A []int) []int {
	max := getMax(A)
	passes := int(math.Log10(float64(max)) + 1)

	for pass := 0; pass < passes; pass++ {
		holding := make([]int, len(A))
		buckets := [10][]int{}
		for i := range A {
			bucket := A[i] / int(math.Pow(10, float64(pass))) % 10
			buckets[bucket] = append(buckets[bucket], A[i])
		}
		k := 0
		for i := range buckets {
			for j := range buckets[i] {
				holding[k] = buckets[i][j]
				k++
			}
		}
		A = holding
	}
	return A
}
