package cci

import "sort"

// Check Permutation: Given two strings,write a method to decide if one is a permutation of the
// other.

// time O(N log n)
// space O(1)
func CheckPermutations(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1sorted := sortString(s1)
	s2sorted := sortString(s2)
	for i := 0; i < len(s1sorted); i++ {
		if s1sorted[i] != s2sorted[i] {
			return false
		}
	}
	return true
}

func sortString(s string) string {
	t := []rune(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}

// time O(N)
// space O(N)

func CheckPermutations2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		c, ok := m[r]
		if !ok {
			return false
		}
		c--
		if c < 0 {
			return false
		}
		m[r] = c
	}
	return true
}

// time O(N)
// space O(1)
// assuming the input is ASCII charactors only

func CheckPermutations3(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	letters := [128]int{}
	for i := range s1 {
		letters[s1[i]]++
	}
	for i := range s2 {
		letters[s2[i]]--
		if letters[s2[i]] < 0 {
			return false
		}
	}
	return true
}
