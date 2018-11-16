package ex704

import (
	"golang.org/x/net/html"
	"io"
	"testing"
)

func TestNewReader(t *testing.T) {
	buffer := make([]byte, 5)
	r := NewReader("1234567")

	c, err := r.Read(buffer)
	if c != 5 {
		t.Errorf("%v != 5", c)
	}
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}

	c, err = r.Read(buffer)
	if c != 2 {
		t.Errorf("%v != 2", c)
	}
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}

	c, err = r.Read(buffer)
	if c != 0 {
		t.Errorf("%v != 0", c)
	}
	if err != io.EOF {
		t.Errorf("%v != %v", err, io.EOF)
	}
}

func TestNewReaderParser(t *testing.T) {
	doc, err := html.Parse(NewReader("<html><head></head><body>test</body></html>"))
	if err != nil {
		t.Errorf("%v", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.Data
	expected := "test"
	if node != expected {
		t.Errorf("%v != %v", node, expected)
	}
}
