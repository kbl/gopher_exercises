package treesort

import (
	"bytes"
	"fmt"
)

func (t *tree) String() string {
	var buffer bytes.Buffer
	buffer.WriteByte('[')
	str(&buffer, t)
	buffer.WriteByte(']')
	return buffer.String()
}

func str(buffer *bytes.Buffer, t *tree) {
	if t == nil {
		return
	}
	str(buffer, t.left)
	if buffer.Len() > len("[") {
		buffer.WriteByte(' ')
	}
	fmt.Fprintf(buffer, "%d", t.value)
	str(buffer, t.right)
}
