package ex710

import "sort"

func IsPalindrome(w sort.Interface) bool {
	half := w.Len() / 2
	for i := 0; i < half; i++ {
		j := w.Len() - i - 1
		if w.Less(i, j) || w.Less(j, i) {
			return false
		}
	}
	return true
}
