package cci

import "sort"

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

// time O(N) for N as the length of the string
// space O(N)
func IsUnique1(s string) bool {
	set := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}

// time O(N log N)
// space O(1)
func IsUnique2(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := []rune(s)
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	t := ss[0]
	for _, r := range ss[1:] {
		if t == r {
			return false
		}
		t = r
	}
	return true
}

// time O(1) since the loop will be on maximun of 128 characters
// space O(1)
// assuming the input is ASCII charactors

func IsUnique3(s string) bool {
	if len(s) > 128 {
		return false
	}
	charSet := [128]bool{}
	for _, c := range s {
		idx := int(c)
		if charSet[idx] {
			return false
		}
		charSet[idx] = true
	}
	return true
}

// O(1)
// space(1)
// assuming input is lower case ASCII charactors
func IsUnique4(s string) bool {
	var checker int32
	for _, c := range s {
		val := int32(c) - int32('a')
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}
