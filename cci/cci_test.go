package cci_test

import (
	"reflect"
	"testing"

	"github.com/Ahmed-Sermani/algos/cci"
	"github.com/Ahmed-Sermani/algos/structures"
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

func TestStringCompression(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{"aaabccccd", "a3b1c4d1"},
		{"aaaaaaaaaab", "a10b1"},
		{"AAAAABBBnn", "A5B3n2"},
		{"John", "John"},
	}
	for _, test := range tests {
		assert.Equal(t, test.expect, cci.StringCompression(test.s))
	}
}

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		in     [][]int
		expect [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{3, 4, 5},
				{6, 7, 8},
			},
			[][]int{
				{6, 3, 1},
				{7, 4, 2},
				{8, 5, 3},
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expect, cci.RotateMatrix(test.in))
	}
}

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		in     [][]int
		expect [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{1, 1, 1},
				{1, 2, 0},
			},
			[][]int{
				{1, 2, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{1, 0, 1},
				{1, 2, 1},
			},
			[][]int{
				{1, 0, 3},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
	}
	for _, test := range tests {
		cp := make([][]int, len(test.in))
		for i := range cp {
			cp[i] = make([]int, len(test.in[i]))
			copy(cp[i], test.in[i])
		}
		assert.Equal(t, test.expect, cci.ZeroMatrix(test.in))
		assert.Equal(t, test.expect, cci.ZeroMatrix1(cp))
	}

}

func TestStringRotation(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		expect bool
	}{
		{"watterbottle", "erbottlewatt", true},
		{"abcd", "dabc", true},
		{"one", "two", false},
		{"333", "4444", false},
	}
	for _, test := range tests {
		assert.Equal(t, test.expect, cci.StringRotation(test.s1, test.s2))
	}
}

func TestRemoteDup(t *testing.T) {
	tests := []struct {
		in, expect *structures.LinkedList[int]
	}{
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 4, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 1, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
	}
	for _, test := range tests {
		cci.RemoveDup(test.in)
		assert.True(t, reflect.DeepEqual(test.in, test.expect))
	}
}

func TestRemoteDup2(t *testing.T) {
	tests := []struct {
		in, expect *structures.LinkedList[int]
	}{
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 4, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 3, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
		{
			in:     buildLinkedListFromSlice([]int{1, 2, 1, 3, 4}),
			expect: buildLinkedListFromSlice([]int{1, 2, 3, 4}),
		},
	}
	for _, test := range tests {
		cci.RemoveDup2(test.in)
		assert.True(t, reflect.DeepEqual(test.in, test.expect))
	}
}

func buildLinkedListFromSlice(arr []int) *structures.LinkedList[int] {
	prevNode := &structures.Node[int]{
		Data: arr[0],
	}
	head := prevNode
	for _, elem := range arr[1:] {
		currNode := &structures.Node[int]{
			Data: elem,
			Prev: prevNode,
		}
		prevNode.Next = currNode
		prevNode = currNode
	}
	return &structures.LinkedList[int]{Head: head}
}
