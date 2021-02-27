package my_expand

// todo handling illegal tokens
func Expand(s string, op func(string) string) string {
	tokenStart, tokenEnd := -1, -1
	previousTokenEnd := 0
	sRunes := []rune(s)
	output := ""
	for i, r := range sRunes {
		if r == '$' {
			tokenStart = i
			continue
		}
		if tokenStart != -1 && r == ' ' {
			tokenEnd = i
			output += string(sRunes[previousTokenEnd:tokenStart])
			output += op(string(sRunes[tokenStart+1 : tokenEnd]))
			previousTokenEnd = tokenEnd
			tokenStart = -1
		}
	}
	if tokenStart != -1 {
		output += string(sRunes[previousTokenEnd:tokenStart])
		output += op(string(sRunes[tokenStart+1:]))
	} else {
		output += string(sRunes[previousTokenEnd:])
	}
	return output
}
