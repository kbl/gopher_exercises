package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
    "strings"
    "regexp"
    "strconv"

    "golang.org/x/net/html"
)

//const url = "https://boardgamegeek.com/geeklist/235895/29-polski-mathandel-polish-math-trade/page/"
const url = "file:///home/mapi/mathandle.html"
var match, _ = regexp.Compile("Average Rating:([0-9.]+)")

type Game struct {
    Id int
    Name string
    Rating float64
    Url string
    Prices []string
}

func main() {
    var allGames []Game
    ch := make(chan []Game)
    pages := 70
    if strings.HasPrefix(url, "file://") {
        pages = 1
    }
    for i := 1; i <= pages; i += 1 {
        customUrl := url
        if ! strings.HasPrefix(url, "file://") {
            customUrl = url + strconv.Itoa(i)
        }
        go parsePage(ch, customUrl)
    }
    for i := 1; i <= pages; i += 1 {
        fmt.Println(i)
        allGames = append(allGames, <-ch...)
    }
    gamesMap := make(map[string][]Game)
    for _, g := range allGames {
        // if gamesMap[g.Name] == nil {
        //     gamesMap[g.Name] = []Game
        // }
        gamesMap[g.Name] = append(gamesMap[g.Name], g)

    }

    fillPrices(gamesMap)

    fmt.Println(len(gamesMap))
    for n, games := range gamesMap {
        fmt.Println(n)
        for _, g := range games {
            fmt.Printf("%d ", g.Id)
        }
        fmt.Println()
        fmt.Println()
    }
    fmt.Println(len(gamesMap))
    fmt.Println(allGames)
}

func parsePage(out chan<- []Game, url string) {
    body, err := getPageBody(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    doc, err := html.Parse(strings.NewReader(string(body)))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // defer func() {
    //     if r := recover(); r != nil {
    //         fmt.Println(url)
    //         fmt.Println("Recovered in f", r)
    //     }
    // }()

    out <- findGames(nil, doc)
}

func findGames(games []Game, node *html.Node) []Game {
    if node.Type == html.ElementNode && node.Data == "dl" {
        for _, a := range node.Attr {
            if a.Key == "id" && strings.HasPrefix(a.Val, "body_listitem") {
                game := parseGameNode(nil, node)
                if game[0].Rating >= 7.45 {
                    games = append(games, game[0])
                }
            }
        }
    }
    for child := node.FirstChild; child != nil; child = child.NextSibling {
        games = findGames(games, child)
    }
    return games
}

func parseGameNode(games []Game, node *html.Node) []Game {
    if node.Type == html.ElementNode && node.Data == "div" {
        for _, a := range node.Attr {
            if a.Key == "class" && a.Val == "fl" {
                content := node.FirstChild

                ordinalTag := content.NextSibling
                id, err := strconv.Atoi(ordinalTag.FirstChild.Data[:len(ordinalTag.FirstChild.Data) - 1])
                if err != nil {
                    fmt.Println(err)
                    os.Exit(1)
                }

                nameTag := ordinalTag.NextSibling.NextSibling
                var url string
                for _, a := range nameTag.Attr {
                    if a.Key == "href" {
                        url = "https://boardgamegeek.com" + a.Val
                    }
                }
                name := nameTag.FirstChild.Data

                ratingTag := nameTag.NextSibling.NextSibling
                var rating float64
                if ratingTag != nil && ratingTag.FirstChild != nil {
                    ratingBody := ratingTag.FirstChild.Data
                    rating, err = strconv.ParseFloat(match.FindStringSubmatch(ratingBody)[1], 32)
                    if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                    }
                }
                return append(games, Game{id, name, rating, url, nil})
            }
        }
    }
    for child := node.FirstChild; child != nil ; child = child.NextSibling {
        games = parseGameNode(games, child)
    }
    return games
}

func getPageBody(url string) ([]byte, error) {
    if strings.HasPrefix(url, "file://") {
        return ioutil.ReadFile(url[len("file://"):])
    }

    response, err := http.Get(url)
    if err != nil {
        return  nil, err
    }

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return body, nil
}

type Price struct {
    Name string
    Prices []string
}

func fillPrices(games map[string][]Game) {
    ch := make(chan Price)
    for _, gs := range games {
        go findPrice(ch, gs[0])
    }

    for _, _ = range games {
        price := <-ch
        fmt.Println(price)
        for _, g := range games[price.Name] {
            g.Prices = price.Prices
        }
    }
}

func findPrice(ch chan<- Price, g Game) {
    body, err := getPageBody(g.Url)
    ioutil.WriteFile("/tmp/dat1", body, 0644)
    fmt.Println(g.Url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    doc, err := html.Parse(strings.NewReader(string(body)))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    prices := findPrices(nil, doc)
    ch<- Price{g.Name, prices}
}

func findPrices(prices []string, node *html.Node) []string {
    if node.Type == html.ElementNode && node.Data == "div" {
        for _, a := range node.Attr {
            // fmt.Println(node)
            if a.Key == "class" && strings.HasPrefix(a.Val, "summary-sale-item-price") {
                fmt.Println("xxx")
                tag := node.FirstChild
                if tag.Data == "span" {
                    tag = tag.NextSibling
                }
                // if tag.Data != "strong" {
                //     fmt.Println("DUPA")
                //     os.Exit(1)
                // }
                prices = append(prices, tag.FirstChild.Data)
            }
        }
    }
    for child := node.FirstChild; child != nil; child = child.NextSibling {
        prices = findPrices(prices, child)
    }
    return prices
}
