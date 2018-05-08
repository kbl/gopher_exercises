package comma

import (
	"bytes"
	"strings"
)

const (
	separator        = ","
	decimalDelimiter = "."
	signs            = "+-"
)

func Comma(s string) string {
	var b bytes.Buffer
	var sign string
	var decimalPart string

	if strings.IndexAny(s, signs) == 0 {
		sign = s[:1]
		s = s[1:]
	}

	n := len(s)

	if n <= 3 {
		return sign + s
	}

	if i := strings.Index(s, decimalDelimiter); i > 0 {
		n = i
		decimalPart = s[n:]
		s = s[:n]
	}

	start := n % 3
	if start == 0 {
		start = 3
	}
	b.WriteString(sign)
	b.WriteString(s[:start])

	for i := start; i < n; i += 3 {
		b.WriteString(separator)
		b.WriteString(s[i : i+3])
	}
	b.WriteString(decimalPart)

	return b.String()
}
