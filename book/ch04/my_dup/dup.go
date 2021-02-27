package my_dup

import "unicode"

func Dup(s []string) []string {
	if len(s) == 0 {
		return s
	}
	c := s[:1]
	prev := s[0]
	for _, v := range s[1:] {
		if v == prev {
			continue
		} else {
			prev = v
			c = append(c, v)
		}
	}
	return c
}

func DupWhitespace(s []byte) []byte {
	ret := s[:0]

	wasWhiteSpace := false
	for _, r := range string(s) {
		if unicode.IsSpace(r) {
			if !wasWhiteSpace {
				ret = append(ret, byte(' '))
			}
			wasWhiteSpace = true
		} else {
			for _, b := range []byte(string(r)) {
				ret = append(ret, b)
			}
			wasWhiteSpace = false
		}
	}

	return ret
}
