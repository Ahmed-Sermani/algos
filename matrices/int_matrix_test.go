package matrices_test

import (
	"reflect"
	"testing"

	"github.com/Ahmed-Sermani/algos/matrices"
)

func TestIntMatrix(t *testing.T) {
	mtx := matrices.NewIntMatrix(3, 3)
	t.Run("Test Init", func(t *testing.T) {
		if len(mtx) != 3 && mtx.RowsLen() != 3 {
			t.Log("Incorrect Init N.rows")
			t.Fail()
		}
		if len(mtx[0]) != 3 && mtx.ColumnsLen() != 0 {
			t.Log("Incorrect Init N.cols")
			t.Fail()
		}
		for i := 0; i < mtx.RowsLen(); i++ {
			for j := 0; j < mtx.ColumnsLen(); j++ {
				if mtx.GetInt(i, j) != 0 {
					t.Logf("element is not inited to zero: (%d, %d)", i, j)
					t.Fail()
				}
			}
		}
	})

	t.Run("Test accessor and mutator", func(t *testing.T) {
		mtx.SetInt(2, 1, 1)
		if mtx.GetInt(1, 1) != 2 {
			t.Fatal()
		}
	})
}

func TestIntMatrixMultiplication(t *testing.T) {
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
	resMtx := matrices.NaiveMultiplication(mtx1, mtx2)
	expected := matrices.IntMatrix{
		{11, 25, 25},
		{27, 54, 52},
		{34, 101, 105},
	}
	if !reflect.DeepEqual(expected, resMtx) {
		t.Fatal("incorrect multiplication")
	}
}

func TestEquality(t *testing.T) {
	tests := []struct {
		mtx1  matrices.IntMatrix
		mtx2  matrices.IntMatrix
		equal bool
	}{
		{
			mtx1: matrices.IntMatrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			mtx2: matrices.IntMatrix{
				{7, 8},
				{9, 10},
				{11, 12},
			},
			equal: false,
		},
		{
			mtx1: matrices.IntMatrix{
				{1, 2},
				{4, 5},
			},
			mtx2: matrices.IntMatrix{
				{7, 8},
				{9, 10},
			},
			equal: false,
		},
		{
			mtx1: matrices.IntMatrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			mtx2: matrices.IntMatrix{
				{1, 2, 3},
				{4, 5, 6},
			},
			equal: true,
		},
	}
	for _, test := range tests {
		res := matrices.Equality(test.mtx1, test.mtx2)
		if test.equal && !res {
			t.Fatal()
		}
		if !test.equal && res {
			t.Fatal()
		}
	}

}
