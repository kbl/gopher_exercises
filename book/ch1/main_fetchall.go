package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "time"
    "book/ch1/fetch"
)

func main() {
    input_file := os.Args[1]
    f, err := os.Open(input_file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
    content, err := ioutil.ReadAll(f)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
    f.Close()
    urls := strings.Split(string(content), "\n")
    urls = filter(urls)

    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }

    start := time.Now()
    responses := fetch.FetchAll(urls)

    for range urls {
        fmt.Println(<-responses)
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func filter(urls []string) []string {
    ret := make([]string, 0)

    for _, v := range urls {
        if len(v) > 1 {
            ret = append(ret, v)
        }
    }

    return ret
}
