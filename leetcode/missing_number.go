package leetcode

// https://leetcode.com/problems/missing-number/
// the key is to use the formula of the sum of n natural numbers
// 1 + 2 + 3 + ..... + (n - 1) + n = (n(n+1)) / 2

func MissingNumber(nums []int) int {
	n := len(nums)
	total := (n * (n + 1)) / 2
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return total - sum
}

// Bit manipulation solution
// we use XOR
// if we XOR same numbers together we will always get zero, hence they cancel each other.
// 000 ^ 000 = 000
// 001 ^ 001 = 000
// 010010 ^ 010010 = 000000
// We XOR the entire array + a counter from 0 to n.
// the result will be the unique number which's the missing number we're looking for.

func MissingNumberXOR(nums []int) int {
	var res int
	for i := range nums {
		res ^= i ^ nums[i]
	}
	res ^= len(nums)
	return res
}
