package count

import (
	"bufio"
	"io"
	"log"
	"unicode"
	"unicode/utf8"
)

type RuneStats struct {
	Runes    map[rune]int
	Lenghts  [utf8.UTFMax + 1]int
	Invalid  int
	Letters  map[rune]int
	Digits   map[rune]int
	Graphics map[rune]int
}

func Runes(input io.Reader) RuneStats {
	ret := RuneStats{
		Runes:    make(map[rune]int),
		Letters:  make(map[rune]int),
		Digits:   make(map[rune]int),
		Graphics: make(map[rune]int),
	}

	in := bufio.NewReader(input)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Runes: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			ret.Invalid++
			continue
		}
		if unicode.IsDigit(r) {
			ret.Digits[r]++
		}
		if unicode.IsLetter(r) {
			ret.Letters[r]++
		}
		if unicode.IsGraphic(r) {
			ret.Graphics[r]++
		}
		ret.Runes[r]++
		ret.Lenghts[n]++
	}

	return ret
}
