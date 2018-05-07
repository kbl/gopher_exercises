package main

import (
	"book/ch01/lissajous"
	"os"
)

func main() {
	lissajous.Draw(5, os.Stdout)
}
