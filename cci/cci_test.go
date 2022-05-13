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

func TestURLify(t *testing.T) {
	tests := []struct {
		s   string
		n   int
		out string
	}{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, cci.URLify(test.s, test.n))
		assert.Equal(t, test.out, cci.URLify1(test.s, test.n))
	}

}

func TestPalindromePermutation(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"Tact Coa", true},
		{"ahmad sermani", false},
		{"taco cat", true},
		{"aab", true},
		{"carerac", true},
		{"aabaa", true},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, cci.PalindromePermutation(test.in))
		assert.Equal(t, test.out, cci.PalindromePermutation1(test.in))
	}
}

func TestOneWay(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		expect bool
	}{
		{"pale", "pales", true},
		{"pale", "ple", true},
		{"pale", "bale", true},
		{"bale", "pake", false},
		{"ahmad", "ahmad", true},
		{"one", "two", false},
	}
	for _, test := range tests {
		assert.Equal(t, test.expect, cci.OneWay(test.s1, test.s2))
		assert.Equal(t, test.expect, cci.OneWay1(test.s1, test.s2))
	}
}
