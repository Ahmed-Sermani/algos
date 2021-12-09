package array

// ResviorSampling given a stream of integers, return a random sample of size k,
// the probability of each data point point to be selected is 1/n.
// the idea is to use reservoir sampling to get a random sample of size k from a stream of size n
func ReservoirSampling(dataCh <-chan int, k int) []int {
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

// ResviorSingleSampling given a stream of integers, return a random sample
func ReservoirSingleSampling(dataCh <-chan int, k int) int {
	return ReservoirSampling(dataCh, 1)[0]
}
