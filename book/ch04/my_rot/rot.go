package rot

import "book/ch04/rev"

func RotWithRev(s []int, by int) {
	if len(s) == 0 || by < 1 {
		return
	}
	by %= len(s)
	rev.Rev(s[by:])
	rev.Rev(s[:by])
	rev.Rev(s)
}

func Rot(s []int, by int) {
	if len(s) == 0 || by < 1 {
		return
	}

	by %= len(s)
	c := make([]int, by)
	copy(c, s[:by])

	var i int

	for ; i < len(s)-by; i++ {
		s[i] = s[(i+by)%len(s)]
	}
	for j, v := range c {
		s[i+j] = v
	}
}
