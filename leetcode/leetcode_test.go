package leetcode_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{0, 1},
			output: 2,
		},
		{
			input:  []int{3, 0, 1},
			output: 2,
		},
		{
			input:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			output: 8,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, leetcode.MissingNumber(test.input))
		assert.Equal(t, test.output, leetcode.MissingNumberXOR(test.input))
	}
}

func TestDisappearingNumbers(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			out: []int{5, 6},
		},
		{
			in:  []int{1, 1},
			out: []int{2},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, leetcode.FindDisappearedNumbers(test.in))
	}
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, leetcode.ClimbStairs(test.in))
	}

}
