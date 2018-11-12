package counter

import "testing"

func TestWordCounter(t *testing.T) {
	var wc WordCounter
	arg := []byte("word d word x")
	expected := 4
	if actual, _ := wc.Write(arg); actual != expected {
		t.Errorf("wc.Write(%v) = %v, expected %v", arg, actual, expected)
	}
}

func TestLineCounter(t *testing.T) {
	var lc LineCounter
	arg := []byte("line\nline\nline")
	expected := 3
	if actual, _ := lc.Write(arg); actual != expected {
		t.Errorf("lc.Write(%v) = %v, expected %v", arg, actual, expected)
	}
}
