package dynamicprogramming

type subsetMemo [][]*bool

var (
	falsePtr = boolptr(false)
	truePtr  = boolptr(true)
)

func SubsetSum(set []int, k int) bool {
	memo := make(subsetMemo, len(set))
	for i := range memo {
		memo[i] = make([]*bool, k+1)
	}
	return subsetSum(set, len(set)-1, k, memo)
}

func subsetSum(set []int, lastindex, k int, memo subsetMemo) bool {

	if k == 0 {
		return true
	}
	if k < 0 || lastindex < 0 {
		return false
	}

	if v := memo[lastindex][k]; v != nil {
		return *v
	}

	for i := 0; i <= lastindex; i++ {
		if subsetSum(set, i-1, k-set[i], memo) {
			memo[lastindex][k] = truePtr
			return true
		}
	}
	memo[lastindex][k] = falsePtr
	return false
}

func boolptr(b bool) *bool {
	return &b
}
