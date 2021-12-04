package array

import "math/rand"

// QuickSelect returns the kth smallest element in the array a with time complexity O(n).
// https://en.wikipedia.org/wiki/Quickselect
// it's a determinstic algorithm (gives the right answer all the time). the non-deterministic part is the partition.
// which makes the algorithm O(n) as expected running time.
// theoretically for better probability of getting the optimal running time is to pick the pivot and evaluate it before the partition.
// but this is not implemented here.
func QuickSelect(a []int, k int) int {
	return quickSelect(a, k)
}

func quickSelect(a []int, k int) int {
	if len(a) == k+1 {
		return a[k]
	}
	l, r := partition(a)
	if k <= len(l) {
		return quickSelect(l, k)
	} else {
		return quickSelect(r, k)
	}
}

func partition(a []int) ([]int, []int) {
	pivot := a[rand.Intn(len(a))]
	var l, r []int
	for _, v := range a {
		if v <= pivot {
			l = append(l, v)
		} else if v > pivot {
			r = append(r, v)
		}
	}
	return l, r
}
