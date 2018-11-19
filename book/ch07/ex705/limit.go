package limit

import (
	"io"
)

type reader struct {
	wrapped io.Reader
	to_read *int64
}

func (r reader) Read(b []byte) (int, error) {
	if *r.to_read == 0 {
		return 0, io.EOF
	}

	if *r.to_read < int64(len(b)) {
		b = b[:*r.to_read]
	}

	count, err := r.wrapped.Read(b)

	*r.to_read -= int64(count)
	if err == io.EOF {
		*r.to_read = 0
	}

	return count, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := reader{r, &n}
	return lr
}
