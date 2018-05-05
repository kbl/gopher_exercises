package main

import (
	"book/ch1/lissajous"
	"os"
)

func main() {
	lissajous.Draw(5, os.Stdout)
}
