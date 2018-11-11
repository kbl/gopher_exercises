package intset

func (s *IntSet) IntersectWith(t *IntSet) {
	for bucket, word := range t.words {
		if bucket >= len(s.words) {
			break
		}
		s.words[bucket] &= word
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for bucket, word := range t.words {
		if bucket >= len(s.words) {
			break
		}
		s.words[bucket] &= ^word
	}
}

func (s *IntSet) SymetricDifferenceWith(t *IntSet) {
	for bucket, word := range t.words {
		if bucket >= len(s.words) {
			s.words = append(s.words, t.words[bucket:]...)
			break
		}
		s.words[bucket] = xor(s.words[bucket], word)
	}
}

func xor(a, b uint64) uint64 {
	return (^(a & b)) & (a | b)
}
