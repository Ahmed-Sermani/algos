package cci

import "strings"

/*
	Assume the method isSubString which checks if one word is a substring of another.
	Given two strings s1 and s2, write a method to check that s2 is a rotation of s1 using
	only one call to isSubString
*/

// time O(N) where N is the length of the string
// space O(N)
func StringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// by concatenating the rotated string to itself
	// we make sure that if it was rotated from the original
	// then the original string should be contained in the new string

	s2 = s2 + s2
	return isSubString(s1, s2)
}

func isSubString(s1, s2 string) bool {
	return strings.Contains(s2, s1)
}
