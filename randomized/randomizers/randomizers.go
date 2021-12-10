package randomizers

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetRandomWithDistribution
// input: array of integers and their frequencies.
// output: return a random integer from the array according the probability that
// corresponds to the frequency.
// example:
// Let following be the given numbers.
//  arr[] = {10, 30, 20, 40}
//  Let following be the frequencies of given numbers.
//	freq[] = {1, 6, 2, 1}
//  The output should be
//	10 with probability 1/10
//	30 with probability 6/10
//	20 with probability 2/10
//	40 with probability 1/10
func GetRandomWithDistribution(arr []int, freq []int) int {
	n := len(arr)
	prefix := make([]int, n)
	// generate prefix sum
	// the accumulated of frequencies will be used to ceil the random number 1-freq[-1]
	// to the nearest frequency.
	prefix[0] = freq[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + freq[i]
	}
	// generate random number
	r := r.Intn(prefix[n-1]) + 1
	// find the index of ceiling of r in prefix.
	// the result index will have distributed probability of being on of the numbers
	// according the frequencies.
	// example:
	// values = [10, 30, 20, 40]
	// freq = [1, 6, 2, 1]
	// prefix = [1, 7, 9, 10]
	// the probability of getting 10 is 1/10 because to get 10
	// we need to get r as 1.
	// the probability of getting 30 is 6/10 because to get 30
	// we need to get r between 2-7.
	c := findCeil(prefix, r, 0, n-1)
	return arr[c]
}

// findCeil returns the index of the first element in the slice that is greater than or equal to the given value.
// If no such element exists, it returns -1.
func findCeil(arr []int, value, low, high int) int {
	for low < high {
		mid := (low + high) / 2
		if value > arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if arr[low] >= value {
		return low
	}
	return -1
}
