package leetcode

import "math"

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// the key here is to get the intendted index to a number (v - 1) and
// make its value as negative. In the end our answer if the values in the array that's positive

func FindDisappearedNumbers(nums []int) []int {
	for i := range nums {
		index := int(math.Abs(float64(nums[i]))) - 1
		nums[index] = -int(math.Abs(float64(nums[index])))
	}
	result := []int{}
	for i := range nums {
		if nums[i] >= 1 {
			result = append(result, i+1)
		}
	}
	return result
}
