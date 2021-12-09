package matrices_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/randomized/matrices"
)

func TestVerifyMultiplication(t *testing.T) {
	mtx1 := matrices.IntMatrix{
		{1, 2, 1},
		{2, 3, 4},
		{7, 5, 3},
	}

	mtx2 := matrices.IntMatrix{
		{1, 8, 9},
		{3, 6, 6},
		{4, 5, 4},
	}
	muli := matrices.IntMatrix{
		{11, 25, 25},
		{27, 54, 52},
		{34, 101, 105},
	}

	if !matrices.IntVerifyMultiplication(mtx1, mtx2, muli) {
		t.Logf("%v * %v != %v", mtx1, mtx2, muli)
	}
}
