package intset

func (s *IntSet) Elems() []int {
	var elems []int
	for bucket, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				elems = append(elems, bucket*64+i)
			}
		}
	}
	return elems
}
