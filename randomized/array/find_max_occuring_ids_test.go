package array_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/randomized/array"
)

func TestFindMaxOccuringIdx(t *testing.T) {
	arr := []int{1, 2, 3, 4, 1}
	if res := array.FindMaxOccuringIdx(arr); res != 0 && res != 4 {
		t.Fail()
	}
}
