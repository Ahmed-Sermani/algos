package sorts

func QuickSort(A []int) []int {
	quicksort(A, 0, len(A)-1)
	return A
}

func quicksort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quicksort(A, p, q-1)
		quicksort(A, q+1, r)
	}
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
