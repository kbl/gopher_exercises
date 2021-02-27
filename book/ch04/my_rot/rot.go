package my_rot

import "github.com/kbl/gopher_exercises/book/ch04/my_rev"

func RotWithRev(s []int, by int) {
	if len(s) == 0 || by < 1 {
		return
	}
	by %= len(s)
	my_rev.Rev(s[by:])
	my_rev.Rev(s[:by])
	my_rev.Rev(s)
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
