package matrices

import (
	"math/rand"
	"time"
)

type Matrix interface {
	RowsLen() int
	ColumnsLen() int
}

type MatrixAccessorInt interface {
	Matrix
	GetInt(int, int) int
}

type MatrixMutatorInt interface {
	Matrix
	// val, i, j
	SetInt(int, int, int)
}

type IntMatrix [][]int

func (m IntMatrix) RowsLen() int {
	return len(m)
}

func (m IntMatrix) ColumnsLen() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m IntMatrix) GetInt(i, j int) int {
	return m[i][j]
}

func (m IntMatrix) SetInt(val, i, j int) {
	m[i][j] = val
}

// initialize new matrix with the specified size
func NewIntMatrix(rows, cols int) IntMatrix {
	var m IntMatrix
	for rows != 0 {
		m = append(m, make([]int, cols))
		rows--
	}
	return m
}

func NaiveMultiplication(A, B MatrixAccessorInt) IntMatrix {
	if A.ColumnsLen() != B.RowsLen() {
		panic("can not multiply")
	}
	resMtx := NewIntMatrix(A.RowsLen(), B.ColumnsLen())
	for i := 0; i < A.RowsLen(); i++ {
		for j := 0; j < B.ColumnsLen(); j++ {
			for k := 0; k < B.ColumnsLen(); k++ {
				resMtx[i][j] += A.GetInt(i, k) * B.GetInt(k, j)
			}
		}
	}
	return resMtx
}

func Equality(A, B MatrixAccessorInt) bool {
	if A.RowsLen() != B.RowsLen() || A.ColumnsLen() != B.ColumnsLen() {
		return false
	}
	n := A.RowsLen()
	m := A.ColumnsLen()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A.GetInt(i, j) != B.GetInt(i, j) {
				return false
			}
		}
	}
	return true
}

func invalidSquaresAndEqualSizeMtx(mtxs ...Matrix) bool {
	rows := mtxs[0].RowsLen()
	cols := mtxs[0].ColumnsLen()
	for _, mtx := range mtxs {
		if mtx.ColumnsLen() != mtx.RowsLen() || mtx.RowsLen() != rows || mtx.ColumnsLen() != cols {
			return true
		}
	}
	return false
}

func genVerticalVector(n int, possibleVals []int) MatrixAccessorInt {
	r := NewIntMatrix(n, 1)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		pick := possibleVals[rand.Int()%len(possibleVals)]
		r[i][0] = pick
	}
	return r
}
