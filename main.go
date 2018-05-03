package main

import (
    "fmt"
    "book/ch1/fetch"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        body, err := fetch.Fetch(url)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", body)
    }
}
