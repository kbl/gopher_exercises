package expand

import (
	"fmt"
	"testing"
)

func duble(s string) string {
	return s + s
}

func TestExpand(t *testing.T) {
	expanded := Expand("AąAa$bc xx", duble)
	if expanded != "AąAabcbc xx" {
		t.Error(fmt.Sprintf(`Expand("AąAa$bc xx", double) != "AąAabcbc xx" (%s)`, expanded))
	}
}

func TestExpandTokenLast(t *testing.T) {
	expanded := Expand("AąAa$bc", duble)
	if expanded != "AąAabcbc" {
		t.Error(fmt.Sprintf(`Expand("AąAa$bc", double) != "AąAabcbc" (%s)`, expanded))
	}
}

func TestExpandMultipleTokens(t *testing.T) {
	expanded := Expand("AąAa$bc $ab", duble)
	if expanded != "AąAabcbc abab" {
		t.Error(fmt.Sprintf(`Expand("AąAa$bc $ab", double) != "AąAabcbc abab" (%s)`, expanded))
	}
}

// func TestExpandIllegalToken(t *testing.T) {
// 	expanded := Expand("x$", duble)
// 	if expanded != "x$" {
// 		t.Error(fmt.Sprintf(`Expand("x$", double) != "x$" (%s)`, expanded))
// 	}
// }
// func TestExpandIllegalToken(t *testing.T) {
// 	expanded := Expand("x$ ", duble)
// 	if expanded != "x$ " {
// 		t.Error(fmt.Sprintf(`Expand("x$ ", double) != "x$ " (%s)`, expanded))
// 	}
// }
// func TestExpandIllegalToken(t *testing.T) {
// 	expanded := Expand("x$ x", duble)
// 	if expanded != "x$ x" {
// 		t.Error(fmt.Sprintf(`Expand("x$ x", double) != "x$ x" (%s)`, expanded))
// 	}
// }
