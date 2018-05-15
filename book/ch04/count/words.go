package count

import (
	"bufio"
	"io"
	"log"
	"unicode/utf8"
)

type WordStats struct {
	Words   map[string]int
	Lenghts map[int]int
}

func Words(input io.Reader) WordStats {
	ret := WordStats{
		Words:   make(map[string]int),
		Lenghts: make(map[int]int),
	}

	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		w := in.Text()
		ret.Words[w]++
		ret.Lenghts[utf8.RuneCountInString(w)]++
	}

	err := in.Err()
	if err != nil {
		log.Fatalf("words: %v\n", err)
	}

	return ret
}
