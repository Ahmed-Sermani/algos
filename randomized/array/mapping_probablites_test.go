package array

import "testing"

func TestMappingProbablities(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := MappingProbablities()
		if a < 1 || a > 8 {
			t.Fatalf("a = %d", a)
		}
	}
}

func TestMappingProbablities2(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := MappingProbablities2()
		if a != 0 && a != 1 {
			t.Fatalf("a = %d", a)
		}
	}
}

func TestShuffle(t *testing.T) {
	old := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := Shuffle(old)
	t.Logf("%v", shuffled)
}
