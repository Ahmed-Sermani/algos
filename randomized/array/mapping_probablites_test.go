package array

import "testing"

func TestMappingProbablities(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := mappingProbablities()
		if a < 1 || a > 8 {
			t.Fatalf("a = %d", a)
		}
	}
}
