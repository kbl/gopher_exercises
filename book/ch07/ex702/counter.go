package counter

import "io"

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func (c countingWriter) Write(buffer []byte) (int, error) {
	count, err := c.writer.Write(buffer)
	*c.count += int64(count)
	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter int64
	wrapped := countingWriter{w, &counter}
	return wrapped, wrapped.count
}
