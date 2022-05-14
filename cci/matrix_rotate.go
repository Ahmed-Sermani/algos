/*
	Give a N*N image. Rotate the image 90 degrees clockwise.
	Can it be done inplace?

	| a | b | c | d |     | m | i | e | a |
	| e | f | g | h | --> | n | j | f | b |
	| i | j | k | l | --> | o | k | g | c |
	| m | n | o | p |     | p | l | h | d |

	We can reach the rotated matrix in two ways:
		1 - Transposing the matrix then inverting the columns.
			* time O(N^2) but it requires two loops.
			* space O(1)
		2 - Performing a series of somewhat complex swaps in one loop.
			* time O(N^2)
			* space O(1)
*/
package cci

func RotateMatrix(m [][]int) [][]int {
	if len(m) == 0 || len(m[0]) != len(m) {
		panic("not N*N matrix")
	}
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < (n / 2); j++ {
			m[i][j], m[i][n-1-j] = m[i][n-1-j], m[i][j]
		}
	}
	return m
}
