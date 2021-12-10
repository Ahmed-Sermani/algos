package randomized

import (
	"math/rand"
	"time"
)

// IsPrime returns true if n is prime.
// Using Fermat's little theorem, we can find a prime number p such that,
// when k = 20 the error = 1/2^20
func IsPrime(n, k int) bool {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	// cornor cases
	if n == 1 || n == 4 {
		return false
	} else if n == 2 || n == 3 {
		return true
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < k; i++ {
		a := r.Intn(n-4) + 2
		if gcd(a, n) != 1 {
			return false
		}
	}
	return true
}
