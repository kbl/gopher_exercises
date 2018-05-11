package rot

import "book/ch04/rev"

func Rot(s []int, by int) {
	if by > len(s) {
		return
	}
	rev.Rev(s[by:])
	rev.Rev(s[:by])
	rev.Rev(s)
}
