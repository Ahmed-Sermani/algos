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
