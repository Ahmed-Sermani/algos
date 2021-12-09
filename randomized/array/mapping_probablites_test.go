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

func TestMappingProbablities2(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := mappingProbablities2()
		if a != 0 && a != 1 {
			t.Fatalf("a = %d", a)
		}
	}
}

func TestShuffle(t *testing.T) {
	old := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := shuffle(old)
	t.Logf("%v", shuffled)
}

func TestResviorSampling(t *testing.T) {
	dataCh := make(chan int, 100)
	go func() {
		for i := 0; i < 10000; i++ {
			dataCh <- i
		}
		close(dataCh)
	}()

	res := reservoirSampling(dataCh, 100)
	t.Logf("%v", res)
}
