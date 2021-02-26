package main

import (
	"github.com/kbl/gopher_exercises/book/ch01/lissajous"
	"os"
)

func main() {
	lissajous.Draw(5, os.Stdout)
}
