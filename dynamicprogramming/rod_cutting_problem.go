package dynamicprogramming

func RobCutting(length int, prices []int) (int, []int) {
	return robCutting(length, prices)
}

// func robCutting(length int, prices []int, memo map[int]int) int {
// 	if v, ok := memo[length]; ok {
// 		return v
// 	}
// 	max := -1
// 	for i := 1; i <= length; i++ {
// 		tmp := prices[i] + robCutting(length-i, prices, memo)
// 		if tmp > max {
// 			max = tmp
// 		}
// 	}
// 	memo[length] = max
// 	return max
// }

func robCutting(length int, prices []int) (int, []int) {
	memo := make([]int, length+1)
	cuts := make([]int, length+1)
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			tmp := prices[j] + memo[i-j]
			if tmp > memo[i] {
				memo[i] = tmp
				cuts[i] = j
			}
		}
	}
	answer := []int{}
	n := length
	for n > 0 {
		answer = append(answer, cuts[n])
		n -= cuts[n]
	}
	return memo[length], answer
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
