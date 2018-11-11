package intset

func (s *IntSet) IntersectWith(t *IntSet) {
	var words []uint64
	for bucket, word := range t.words {
		if bucket > len(s.words) {
			break
		}
		words = append(words, word&s.words[bucket])
	}
	s.words = words
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	var words []uint64
	for bucket, word := range s.words {
		if bucket > len(s.words) {
			break
		}
		words = append(words, word&s.words[bucket])
	}
	s.words = words
}

func (s *IntSet) SymetricDifferenceWith(t *IntSet) {
}
