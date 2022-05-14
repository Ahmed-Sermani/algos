/*
	Give M*N matrix, if any of its sells are 0 then the entire row can column are also set to 0.
*/
package cci

// time O(NM)
// space O(z) where z is the number of found zeros in the matrix
func ZeroMatrix(m [][]int) [][]int {
	if len(m) == 0 || len(m[0]) == 0 {
		panic("invalid input")
	}
	locations := [][2]int{}
	for i := 0; i < len(m); i++ {
		for k := 0; k < len(m[0]); k++ {
			if p := m[i][k]; p == 0 {
				locations = append(locations, [2]int{i, k})
			}
		}
	}
	for _, loc := range locations {
		row := loc[0]
		col := loc[1]
		for i := 0; i < len(m); i++ {
			m[row][i] = 0
		}
		for i := 0; i < len(m[0]); i++ {
			m[i][col] = 0
		}
	}
	return m
}

// time O(NM)
// space O(1)
func ZeroMatrix1(m [][]int) [][]int {
	if len(m) == 0 || len(m[0]) == 0 {
		panic("invalid input")
	}
	var rowHasZero, colHasZero bool
	for i := 0; i < len(m[0]); i++ {
		if m[0][i] == 0 {
			rowHasZero = true
		}
	}
	for i := 0; i < len(m); i++ {
		if m[i][0] == 0 {
			colHasZero = true
		}
	}
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				m[i][0] = 0
				m[0][j] = 0
			}
		}
	}
	for i := 0; i < len(m); i++ {
		if m[i][0] == 0 {
			nullifyRow(m, i)
		}
	}
	for i := 0; i < len(m[0]); i++ {
		if m[0][i] == 0 {
			nullifyCol(m, i)
		}
	}
	if rowHasZero {
		nullifyRow(m, 0)
	}
	if colHasZero {
		nullifyCol(m, 0)
	}
	return m
}

func nullifyCol(m [][]int, col int) {
	for i := 0; i < len(m); i++ {
		m[i][col] = 0
	}
}

func nullifyRow(m [][]int, row int) {
	for i := 0; i < len(m[0]); i++ {
		m[row][i] = 0
	}
}
