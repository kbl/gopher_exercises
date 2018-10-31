package ex516

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	parts := []string{"1", "22", "333"}
	separator := ","
	result := Join(separator, parts...)

	if result != strings.Join(parts, separator) {
		t.Errorf("Join(%v, %v) = %v, wants %v", separator, parts, result, strings.Join(parts, separator))
	}
}
