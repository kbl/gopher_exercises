package comma

import (
	"testing"
)

func TestCommaRecursive(t *testing.T) {
	testCases := map[string]string{
		"":        "",
		"123":     "123",
		"1234":    "1,234",
		"123456":  "123,456",
		"1234567": "1,234,567",
	}

	for testData, expected := range testCases {
		result := Comma(testData)
		if result != expected {
			t.Errorf("Comma(%q) == %q, wants %q", testData, result, expected)
		}
	}
}

func TestCommaRecursiveOptionalSign(t *testing.T) {
	testCases := map[string]string{
		"-123":     "-123",
		"+1234":    "+1,234",
		"-123456":  "-123,456",
		"+1234567": "+1,234,567",
	}

	for testData, expected := range testCases {
		result := Comma(testData)
		if result != expected {
			t.Errorf("Comma(%q) == %q, wants %q", testData, result, expected)
		}
	}
}

func TestCommaRecursiveFloatingPoint(t *testing.T) {
	testCases := map[string]string{
		"123.4":        "123.4",
		"123.45":       "123.45",
		"1234.5":       "1,234.5",
		"1234.5678":    "1,234.5678",
		"123456.7":     "123,456.7",
		"1234567.8":    "1,234,567.8",
		"1234567.8901": "1,234,567.8901",
	}

	for testData, expected := range testCases {
		result := Comma(testData)
		if result != expected {
			t.Errorf("Comma(%q) == %q, wants %q", testData, result, expected)
		}
	}
}

func TestCommaRecursiveFloatingPointWithSign(t *testing.T) {
	testCases := map[string]string{
		"+123.4":        "+123.4",
		"+123.45":       "+123.45",
		"+1234.5":       "+1,234.5",
		"-1234.5678":    "-1,234.5678",
		"-123456.7":     "-123,456.7",
		"-1234567.8":    "-1,234,567.8",
		"-1234567.8901": "-1,234,567.8901",
	}

	for testData, expected := range testCases {
		result := Comma(testData)
		if result != expected {
			t.Errorf("Comma(%q) == %q, wants %q", testData, result, expected)
		}
	}
}
