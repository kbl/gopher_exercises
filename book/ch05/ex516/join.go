package ex516

func Join(sep string, parts ...string) string {
	var result string
	lastIndex := len(parts) - 1
	for i, e := range parts {
		result = result + e
		if i < lastIndex {
			result = result + sep
		}
	}
	return result
}
