package cci

import "unicode"

// Palindrome Permutation: Given a string,
// write a function to check if it is a permutation of a palin­drome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters.
// The palindrome does not need to be limited to just dictionary words

// time O(N) where N is the length of the string
// space O(N)
func PalindromePermutation(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		if c == ' ' {
			continue
		}
		m[unicode.ToLower(c)] += 1
	}
	misMatchCount := 0
	for k := range m {
		misMatchCount += m[k] & 1
	}

	return misMatchCount <= 1
}

// time O(N)
// space O(1)
// assuming ASCII input
func PalindromePermutation1(s string) bool {
	m := [128]int{}
	for _, r := range s {
		if r == ' ' {
			continue
		}
		idx := int(unicode.ToLower(r))
		m[idx]++
	}
	misMatchCount := 0
	for i := range m {
		misMatchCount += m[i] & 1
	}
	return misMatchCount <= 1
}
