package matrices

// IntVerifyMultiplication verify that AB = C using Freivalds' algorithm
// https://en.wikipedia.org/wiki/Freivalds%27_algorithm
func IntVerifyMultiplication(A, B, C MatrixAccessorInt) bool {
	if invalidSquaresAndEqualSizeMtx(A, B) {
		panic("matrices not square or not equal sizes: all matrices has to be n x n size")
	}
	n := A.RowsLen()
	r := genVerticalVector(n, []int{0, 1})

	return Equality(NaiveMultiplication(A, NaiveMultiplication(B, r)), NaiveMultiplication(C, r))

}
