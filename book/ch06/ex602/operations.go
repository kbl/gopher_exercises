package intset

func (s *IntSet) AddAll(x ...int) {
	for _, e := range x {
		s.Add(e)
	}
}
