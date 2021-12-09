package array

import "testing"

func TestResviorSampling(t *testing.T) {
	dataCh := make(chan int, 100)
	go func() {
		for i := 0; i < 10000; i++ {
			dataCh <- i
		}
		close(dataCh)
	}()

	res := ReservoirSampling(dataCh, 100)
	t.Logf("%v", res)
}
