package anagram

import "reflect"

func AreAnagrams(s1, s2 string) bool {
	s1m := make(map[rune]int)
	s2m := make(map[rune]int)

	for _, r := range s1 {
		s1m[r] += 1
	}

	for _, r := range s2 {
		s2m[r] += 1
	}

	return reflect.DeepEqual(s1m, s2m)
}
