package randomized_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/randomized"
)

func TestIsPrime(t *testing.T) {
	t.Parallel()
	primes := []int{7727, 7741, 7753, 7757, 7759, 7789, 7793, 7817, 7823, 7829, 7841, 7853, 7867, 7873, 7877, 7879, 7883, 7901, 7907, 7919}
	for _, prime := range primes {
		if !randomized.IsPrime(prime, 20) {
			t.Errorf("%d is prime", prime)
		}
	}
	composites := []int{7728, 7742, 7754, 7758, 7760, 7790, 7794, 7818, 7824, 7830, 7842, 7854, 7868, 7874, 7878, 7880, 7884, 7902, 7908, 7920, 999, 1001}
	for _, composite := range composites {
		if randomized.IsPrime(composite, 20) {
			t.Errorf("%d is composite", composite)
		}
	}
}
