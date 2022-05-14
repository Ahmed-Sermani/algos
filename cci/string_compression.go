package cci

import (
	"fmt"
	"strings"
)

/*
	Implement a string compression algorithm that compress repeated characters into numbers.
	If the compressed version wasn't smaller than the original, the method should return the original string.
	Assume the input consists only from upper and lower case ASCII characters.
	Example:
		aaabccccd -> a3b1c4d1
		aaaaaaaaaab -> a10b1
		John -> John
*/

// time O(N) where N is the length of the string
// space O(N) where N is compressed length of the string
func StringCompression(s string) string {
	if len(s) == 0 {
		return s
	}
	var (
		sb        strings.Builder
		lastSeen  byte
		seenCount int
	)
	for i := range s {
		if lastSeen == s[i] {
			seenCount++
		} else {
			if lastSeen == 0 {
				lastSeen = s[i]
				seenCount = 1
				continue
			}
			sb.WriteByte(lastSeen)
			sb.WriteString(fmt.Sprint(seenCount))
			lastSeen = s[i]
			seenCount = 1
		}
	}
	sb.WriteByte(lastSeen)
	sb.WriteString(fmt.Sprint(seenCount))

	compressed := sb.String()
	if len(compressed) > len(s) {
		return s
	}
	return compressed

}
