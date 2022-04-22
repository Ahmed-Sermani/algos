package cci_test

import (
	"testing"

	"github.com/Ahmed-Sermani/algos/cci"
	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"isbe", true},
		{"iiasd", false},
		{"asdfad", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, cci.IsUnique1(test.input))
		assert.Equal(t, test.expected, cci.IsUnique2(test.input))
		assert.Equal(t, test.expected, cci.IsUnique3(test.input))
		assert.Equal(t, test.expected, cci.IsUnique4(test.input))
	}
}

func TestCheckPermutations(t *testing.T) {
	tests := []struct {
		input    [2]string
		expected bool
	}{
		{
			[2]string{"ahmad", "maahd"}, true,
		},
		{
			[2]string{"443322", "332244"}, true,
		},
		{
			[2]string{"aasss", "aass"}, false,
		},
		{
			[2]string{"03", "12"}, false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, cci.CheckPermutations(test.input[0], test.input[1]))
		assert.Equal(t, test.expected, cci.CheckPermutations2(test.input[0], test.input[1]))
		assert.Equal(t, test.expected, cci.CheckPermutations3(test.input[0], test.input[1]))
	}
}
