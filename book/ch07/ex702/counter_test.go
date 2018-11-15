package counter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var b bytes.Buffer
	w, count := CountingWriter(&b)
	fmt.Fprintf(w, "12345")
	fmt.Fprintf(w, "67890")
	if *count != 10 {
		t.Errorf("%v != 10", *count)
	}
}
