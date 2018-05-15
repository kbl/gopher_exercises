package charcount

import (
	"bufio"
	"io"
	"log"
	"unicode"
	"unicode/utf8"
)

type Stats struct {
	Runes   map[rune]int
	Lenghts [utf8.UTFMax + 1]int
	Invalid int
}

func Count(input io.Reader) Stats {
	ret := Stats{Runes: make(map[rune]int)}

	in := bufio.NewReader(input)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Count: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			ret.Invalid++
			continue
		}
		ret.Runes[r]++
		ret.Lenghts[n]++
	}

	return ret
}
