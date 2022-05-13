package cci

/*
	giving three operations that can be done on a string. delete a charactor, add a charactor, or replace a charactor.
 	write a function that takes two strings and return wether or not they're one (or zero) operation away.
	pale -> pale return true (zero op)
	pales -> pale return true (insert or delete)
	spale -> pale
	bake -> pake -> return true (replace)
	bake -> pale -> return false

	difflen = |len(a) - len(b)|
	if difflen > 1: return false
	if difflen == 1: # insert or update might be possible
		n = min(len(a), len(b))
		if the two strings have n shared charactors in order:
			return true
		else return false
	# replace might be possible
	n = len(a) - 1
	if two strings has n shared charactors in order:
		return true
	else return false

	time O(N + M)
	space(min(N, M))
*/

// time O(N + M)
// space O(min(N, M))
func OneWay(s1, s2 string) bool {
	difflen := len(s1) - len(s2)
	if difflen < 0 {
		difflen = -difflen
	}
	switch {
	case difflen > 1:
		return false
	case difflen == 1:
		m := map[rune]struct{}{}
		var mins, maxs string
		if len(s1) > len(s2) {
			mins = s2
			maxs = s1
		} else {
			mins = s1
			maxs = s2
		}
		for _, r := range mins {
			m[r] = struct{}{}
		}
		var gotFirstMiss bool
		for _, r := range maxs {
			if _, found := m[r]; !found {
				if gotFirstMiss {
					return false
				}
				gotFirstMiss = true
			}
		}
	default:
		if s1 == s2 {
			return true
		}
		m := map[rune]struct{}{}
		for _, r := range s1 {
			m[r] = struct{}{}
		}
		var gotFirstMiss bool
		for _, r := range s2 {
			if _, found := m[r]; !found {
				if gotFirstMiss {
					return false
				}
				gotFirstMiss = true
			}
		}
	}
	return true
}

// time O(min(N, M))
// space O(1)
func OneWay1(first, second string) bool {
	difflen := len(first) - len(second)
	if difflen < 0 {
		difflen = -difflen
	}
	if difflen > 1 {
		return false
	}
	var s1, s2 []rune
	if len(first) > len(second) {
		s1 = []rune(second)
		s2 = []rune(first)
	} else {
		s1 = []rune(first)
		s2 = []rune(second)
	}

	var i, j int
	var sawDiff bool
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			if sawDiff {
				return false
			}
			// replace case
			if len(s1) == len(s2) {
				i++
			}
			sawDiff = true
		} else {
			i++
		}
		j++
	}
	return true
}
