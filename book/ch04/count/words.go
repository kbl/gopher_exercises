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
	for {
		if in.Scan() {
			w := in.Text()
			ret.Words[w]++
			ret.Lenghts[utf8.RuneCountInString(w)]++
			continue
		}
		err := in.Err()
		if err == nil {
			break
		}
		log.Fatalf("words: %v\n", err)
	}

	return ret
}
