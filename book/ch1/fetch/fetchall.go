package fetch

import (
    "os"
    "fmt"
    "time"
    "strings"
    "bufio"
)

func FetchAll(urls []string) <-chan string {
    ch := make(chan string)

    for _, url := range urls {
        go fetchAndWrite(url, ch)
    }
    return ch
}

func fetchAndWrite(url string, ch chan<- string) {
    start := time.Now()
    body, err := Fetch(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    domain := getDomain(url)
    r := strings.NewReplacer(".", "_")
    fileName := r.Replace(domain) + start.Format("_2006-01-02T15:04:05")
    f, err := os.Create(fileName)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }
    defer f.Close()

    writer := bufio.NewWriter(f)
    nbytes, err := writer.WriteString(body)

    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func getDomain(url string) string {
    var domain string
    if strings.HasPrefix(url, "http://") {
        domain = strings.Split(url, "http://")[1]
    } else if strings.HasPrefix(url, "https://") {
        domain = strings.Split(url, "https://")[1]
    } else {
        domain = url
    }

    fmt.Println(domain)

    return strings.Split(domain, "/")[0]
}
