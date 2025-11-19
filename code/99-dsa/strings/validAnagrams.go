package strings

import (
	"unicode/utf8"
)

func ValidAnagrams(s1 string, s2 string) bool {

	if c1, c2 := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2); c1 == c2 {
		//method 1
		// create fMap for s1 T&S : O(n)
		// create fMap for s2 T&S : O(n)
		// fMap1 == fMap2 ? true : false O(n)
		size := c1
		fMap1 := make(map[rune]int, size)
		fMap2 := make(map[rune]int, size)

		for _, v := range s1 {
			fMap1[v] += 1
		}

		for _, v := range s2 {
			fMap2[v] += 1
		}
		isAnagram := true
		for key := range fMap1 {
			if fMap1[key] != fMap2[key] {
				isAnagram = false
				break
			}
		}
		return isAnagram

	}
	return false
}
