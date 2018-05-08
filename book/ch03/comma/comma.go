package comma

import "bytes"

func Comma(s string) string {
	var b bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	start := n % 3
	if start == 0 {
		start = 3
	}
	b.WriteString(s[:start])

	for i := start; i < n; i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}

	return b.String()
}
