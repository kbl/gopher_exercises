package ex710

import "testing"

type sortString struct {
	str []rune
}

func ss(str string) *sortString {
	ss := new(sortString)
	ss.str = []rune(str)
	return ss
}

func (s *sortString) String() string {
	return string(s.str)
}

func (s *sortString) Len() int {
	return len(s.str)
}

func (s *sortString) Swap(i, j int) {
	s.str[i], s.str[j] = s.str[j], s.str[i]
}

func (s *sortString) Less(i, j int) bool {
	return s.str[i] < s.str[j]
}

func TestIsPalindrome(t *testing.T) {
	palindromes := []*sortString{ss("abccba"), ss("abcba")}
	notPalindromes := []*sortString{ss("kobyla ma maly bok"), ss("kajaki")}

	for _, w := range palindromes {
		if !IsPalindrome(w) {
			t.Errorf("IsPalindrome(%v) != true", w)
		}
	}
	for _, w := range notPalindromes {
		if IsPalindrome(w) {
			t.Errorf("IsPalindrome(%v) != false", w)
		}
	}
}
