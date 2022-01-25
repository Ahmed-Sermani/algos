package array

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

// Find the max occuring int and return random on of its indices in equal probability
func FindMaxOccuringIdx(arr []int) int {
	freq := make(map[int][]int)
	if len(arr) == 0 {
		panic("arr is empty")
	}
	curMax := -1
	for i := range arr {
		freq[arr[i]] = append(freq[arr[i]], i)
		if len(freq[arr[i]]) > curMax {
			curMax = arr[i]
		}
	}
	maxOccuringIdxs := freq[curMax]
	return maxOccuringIdxs[r.Intn(len(maxOccuringIdxs))]

}
