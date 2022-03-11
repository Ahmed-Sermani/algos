package dynamicprogramming

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRodCutting(t *testing.T) {
	tests := []struct {
		length   int
		prices   []int
		expected int
	}{
		{
			length:   10,
			prices:   []int{0, 5, 8, 9, 10, 17, 17, 20, 24, 30, 31},
			expected: 50,
		},
		{
			length:   5,
			prices:   []int{0, 1, 8, 9, 10, 17},
			expected: 17,
		},
	}
	for _, test := range tests {
		total, cuts := RobCutting(test.length, test.prices)
		fmt.Print(cuts)
		assert.Equal(t, test.expected, total)
	}
}

func TestSubsetSum(t *testing.T) {
	tests := []struct {
		set      []int
		expected bool
		k        int
	}{
		{
			set:      []int{1, 2, 3, 1},
			k:        7,
			expected: true,
		},
		{
			set:      []int{3, 34, 4, 12, 5, 2},
			k:        9,
			expected: true,
		},
		{
			set:      []int{3, 34, 4, 12, 5, 2},
			k:        30,
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, SubsetSum(test.set, test.k))
	}
}

func BenchmarkRobCutting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RobCutting(10, []int{0, 5, 8, 9, 10, 17, 17, 20, 24, 30, 31})
	}
}

func BenchmarkSubsetSum(b *testing.B) {
	set := []int{3, 34, 4, 12, 5, 2, 2, 3, 4, 5, 6, 2, 1, 2, 3, 0, 11, 99, 100, 111}
	k := 9
	memo := make(subsetMemo, len(set))
	for i := range memo {
		memo[i] = make([]*bool, k+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		subsetSum(set, len(set)-1, k, memo)
	}
}
