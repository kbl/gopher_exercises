package limit

import (
	"io"
	"strings"
	"testing"
)

func TestLimit(t *testing.T) {
	buffer := make([]byte, 3)
	reader := strings.NewReader("1234567")
	limited_reader := LimitReader(reader, 5)

	c, err := limited_reader.Read(buffer)
	if c != 3 {
		t.Errorf("%v != 3", c)
	}
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}

	c, err = limited_reader.Read(buffer)
	if c != 2 {
		t.Errorf("%v != 2", c)
	}
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}

	c, err = limited_reader.Read(buffer)
	if c != 0 {
		t.Errorf("%v != 0", c)
	}
	if err != io.EOF {
		t.Errorf("%v != %v", err, io.EOF)
	}

}
