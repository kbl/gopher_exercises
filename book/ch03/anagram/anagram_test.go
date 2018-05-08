package anagram

import (
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	testCases := map[[2]string]bool{
		{"abc", "abc"}:     true,
		{"bac", "abc"}:     true,
		{"cac", "abc"}:     false,
		{"caccb", "cacbc"}: true,
	}

	for testData, expected := range testCases {
		result := AreAnagrams(testData[0], testData[1])
		if result != expected {
			t.Errorf("AreAnagrams(%q, %q) != %v", testData[0], testData[1], expected)
		}
	}
}
