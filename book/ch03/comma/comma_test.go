package comma

import (
	"testing"
)

func TestCommaRecursive(t *testing.T) {
	testCases := map[string]string{
		"":        "",
		"123":     "123",
		"1234":    "1,234",
		"1234567": "1,234,567",
	}

	for testData, expected := range testCases {
		result := Comma(testData)
		if result != expected {
			t.Errorf("Comma(%q) == %q, wants %q", testData, result, expected)
		}
	}
}
