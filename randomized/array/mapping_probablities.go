package array

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// given a function foo that returns 1-3 with equal probability,
// write a function that returns 1-8 with equal probability
// the idea is to we want a way to map numbers of 1-3 to numbers 1-(any mulipable of 8 "8, 16, 24") to have
// equal probability of 1-8 to be returned when modding by 8
func mappingProbablities() int {
	// this gets numbers 1-18
	i := 6*foo() + foo() - 3
	// but we want 1-16
	// if i > 16, recursively call mappingProbablities()
	if i < 17 {
		return (i % 8) + 1
	}
	return mappingProbablities()
}
func foo() int {
	return 1 + r.Intn(3)
}

// give a function foo2 that return 0 with 60% probability and 1 with 40% probability.
// write a function that returns 1 and 0 with equal probability.
// the solution is to call foo2 twice and add the results together.
// in case of (0, 1) and (1, 0) the probability is equal 0.4 * 0.6 = 0.24.
// in case of (0, 0) and (1, 1) recur.
// http://en.wikipedia.org/wiki/Fair_coin#Fair_results_from_a_biased_coin
func mappingProbablities2() int {
	// this gets numbers 0-1
	i, j := foo2(), foo2()
	if i^j == 1 {
		return i
	}
	return mappingProbablities2()
}

func foo2() int {
	d := [...]int{1, 1, 0, 1, 0, 1, 0, 0, 1, 1}
	return d[r.Intn(len(d))]
}

// given a slice of intergers shuffle the slice in place with O(n) time complexity
func shuffle(slice []int) []int {
	slc := make([]int, len(slice))
	copy(slc, slice)
	for i := len(slc) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
	return slc
}

// resviorSampling given a stream of integers, return a random sample of size k
// the idea is to use reservoir sampling to get a random sample of size k from a stream of size n
func reservoirSampling(dataCh <-chan int, k int) []int {
	// create a reservoir
	var n int
	reservoir := make([]int, k)
	for dataPoint := range dataCh {
		if n < k {
			reservoir[n] = dataPoint
		} else {
			j := r.Intn(n + 1)
			if j < k {
				reservoir[j] = dataPoint
			}
		}
		n++
	}
	return reservoir
}
