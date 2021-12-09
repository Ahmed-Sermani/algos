package array

import "math/rand"

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
	return 1 + rand.Intn(3)
}
