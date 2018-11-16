package ex704

import (
	"io"
)

type reader struct {
	buffer  []byte
	pointer int
}

func (r *reader) Read(p []byte) (n int, err error) {
	count := copy(p, r.buffer[r.pointer:])
	r.pointer += count
	return count, nil
}

func NewReader(str string) io.Reader {
	r := new(reader)
	r.buffer = []byte(str)
	return r
}
