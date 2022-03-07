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

func BenchmarkRobCutting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RobCutting(10, []int{0, 5, 8, 9, 10, 17, 17, 20, 24, 30, 31})
	}
}
