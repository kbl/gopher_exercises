package main

import (
    "book/ch3/surface"
    "os"
)

func main() {
    writeToStdout()
}

func writeToStdout() {
    surface.Write(os.Stdout, surface.Saddle)
}
