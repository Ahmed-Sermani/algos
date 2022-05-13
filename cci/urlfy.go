package cci

import "unicode"

// URLify: Write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficient space at the end to hold the additional characters,
// and that you are given the "true" length of the string.

// time O(sN^3) where N is the lenghth of the string and s is the number of spaces.
// space O(1)
func URLify(s string, n int) string {
	ss := []rune(s)
	for i := range ss {
		if ss[i] == ' ' {
			ss = append(ss[:i], append([]rune{'%', '2', '0'}, ss[i+1:]...)...)
			i += 2
		}
	}
	for i := len(ss) - 1; i > 0; i-- {
		if unicode.IsLetter(ss[i]) {
			ss = ss[:i+1]
			break
		}
	}
	return string(ss)
}

// time O(N)
// space O(1)
func URLify1(s string, n int) string {
	ss := []rune(s)
	sp := len(ss) - 1
	sawFirstChar := false
	for i := len(ss) - 1; i > 0; i-- {
		switch {
		case unicode.IsLetter(ss[i]):
			sawFirstChar = true
			ss[i], ss[sp] = ss[sp], ss[i]
			sp--
		case ss[i] == ' ' && sawFirstChar:
			ss[sp] = '0'
			ss[sp-1] = '2'
			ss[sp-2] = '%'
			sp -= 3
		}
	}
	for i := range ss {
		if unicode.IsLetter(ss[i]) {
			return string(ss[i:])
		}
	}
	return s
}
