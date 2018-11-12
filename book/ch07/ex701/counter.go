package counter

import "bufio"

type WordCounter int

type LineCounter int

func (counter *WordCounter) Write(buffer []byte) (int, error) {
	*counter += WordCounter(write(buffer, bufio.ScanWords))
	return int(*counter), nil
}

func (counter *LineCounter) Write(buffer []byte) (int, error) {
	*counter += LineCounter(write(buffer, bufio.ScanLines))
	return int(*counter), nil
}

func write(buffer []byte, scanFunc func([]byte, bool) (int, []byte, error)) int {
	count := 0
	if len(buffer) > 0 {
		count += 1
	}
	for true {
		index, token, _ := scanFunc(buffer, false)
		if len(token) > 0 {
			count += 1
		}
		if index == 0 {
			break
		}
		buffer = buffer[index:]
	}
	return count
}
