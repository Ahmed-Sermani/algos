package leetcode

// dynamic programming problem
// https://leetcode.com/problems/climbing-stairs/

func ClimbStairs(n int) int {
	i, j := 0, 1
	for k := 1; k <= n; k++ {
		curr := i + j
		i = j
		j = curr
	}
	return j
}
