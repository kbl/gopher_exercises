package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const geekList = "https://www.boardgamegeek.com/xmlapi/geeklist/235895"
const geekListFile = "/home/mapi/Desktop/235895.xml"

const gameTemplate = "https://www.boardgamegeek.com/xmlapi/boardgame/%d?pricehistory=1&stats=1"
const gameFileTemplate = "/home/mapi/work/20171103_go/src/mathandel/gry/%d.xml"
const gameTemplateFile = "/home/mapi/Desktop/126000.xml"

type Geeklist struct {
	XMLName xml.Name `xml:"geeklist"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Ordinal     int    `xml:"id,attr"`
	GameId      int    `xml:"objectid,attr"`
	Name        string `xml:"objectname,attr"`
	Description string `xml:"body"`
}

type Boardgames struct {
	XMLName   xml.Name  `xml:"boardgames"`
	Boardgame Boardgame `xml:"boardgame"`
}

type Boardgame struct {
	Year       int       `xml:"yearpublished"`
	Awards     []string  `xml:"boardgamehonor"`
	Categories []string  `xml:"boardgamecategory"`
	Mechanics  []string  `xml:"boardgamemechanic"`
	Rating     float64   `xml:"statistics>ratings>average"`
	Geekrating float64   `xml:"statistics>ratings>bayesaverage"`
	Listings   []Listing `xml:"marketplacehistory>listing"`
}

type Listing struct {
	Body string `xml:",innerxml"`
}

// ---

type GameMetadata struct {
	Id         int
	Year       int
	Awards     int
	Mechanics  []string
	Categories []string
	Rating     float64
	PriceUSD   float64
	PriceEUR   float64
}

var gamesToIgnore = map[string]bool{
	"7 Wonders":                  true,
	"Agricola (revised edition)": true,
	"Agricola":                   true,
	//    "Brass: Lancashire": true,
	"Carcassonne: Winter Edition":                                             true,
	"Dixit Quest":                                                             true,
	"Dominion":                                                                true,
	"Dominion: Cornucopia":                                                    true,
	"Dominion: Empires":                                                       true,
	"Dominion: Intrigue":                                                      true,
	"Dominion: Prosperity":                                                    true,
	"Dominion: Seaside":                                                       true,
	"Pandemic":                                                                true,
	"Patchwork":                                                               true,
	"Power Grid":                                                              true,
	"Puerto Rico + Expansion I":                                               true,
	"Puerto Rico":                                                             true,
	"Small World: Sky Islands":                                                true,
	"Star Trek: Frontiers":                                                    true,
	"Star Wars: Imperial Assault – Alliance Smuggler Ally Pack":               true,
	"Star Wars: Imperial Assault – General Sorin Villain Pack":                true,
	"Star Wars: Rebellion":                                                    true,
	"Star Wars: X-Wing Miniatures Game – Lambda-class Shuttle Expansion Pack": true,
	"Star Wars: X-Wing Miniatures Game – T-70 X-Wing Expansion Pack":          true,
	"Terra Mystica":                             true,
	"Tzolk'in: The Mayan Calendar":              true,
	"War of the Ring (first edition)":           true,
	"War of the Ring: Battles of the Third Age": true,
}

func main() {
	r := fetchUrl(geekListFile)
	var geeklist Geeklist
	xml.Unmarshal(r, &geeklist)

	results := make(chan GameMetadata)
	jobs := make(chan int)

	gameNames := make(map[int]string)
	gameOccurences := make(map[int][]string)
	gameMeta := make(map[int]GameMetadata)

	replacer := strings.NewReplacer("?", "", ",", "", "!", "", " ", "", ")", "", "(", "", ".", "", "–", "", "&", "", "'", "", "/", "", "-", "", ":", "")

	for i, g := range geeklist.Items {
		if _, ok := gamesToIgnore[g.Name]; ok {
			continue
		}
		gameNames[g.GameId] = strings.ToUpper(replacer.Replace(g.Name))
		gameOccurences[g.GameId] = append(gameOccurences[g.GameId], strconv.Itoa(i+1))
	}

	for w := 1; w < 100; w++ {
		go worker(w, jobs, results)
	}
	go submitAllJobs(jobs, gameNames)

	var howMany int
	for _, _ = range gameNames {
		howMany++
		meta := <-results
		gameMeta[meta.Id] = meta
	}

	worthNoting := make(map[int]bool)

	gamesToCheck := make(map[int]bool)
	gamesAbove35 := make(map[int]bool)
	gamesAbove50 := make(map[int]bool)
	gamesAbove60 := make(map[int]bool)

	var allGames []string
	var above35Games []string
	var above50Games []string
	var above60Games []string

	for id, meta := range gameMeta {
		if meta.Rating >= 7.45 {
			worthNoting[id] = true
			fmt.Printf("(greengoo71) %%%s : %s\n", gameNames[id], strings.Join(gameOccurences[id], " "))
			allGames = append(allGames, fmt.Sprintf("%%%s", gameNames[id]))

			if meta.PriceEUR == 0 {
				gamesToCheck[id] = true
			}
			if meta.PriceEUR >= 35 {
				gamesAbove35[id] = true
				above35Games = append(above35Games, fmt.Sprintf("%%%s", gameNames[id]))
			}
			if meta.PriceEUR >= 50 {
				gamesAbove50[id] = true
				above50Games = append(above50Games, fmt.Sprintf("%%%s", gameNames[id]))
			}
			if meta.PriceEUR >= 60 {
				gamesAbove60[id] = true
				above60Games = append(above60Games, fmt.Sprintf("%%%s", gameNames[id]))
			}
		}
	}

	sort.Strings(allGames)
	sort.Strings(above35Games)
	sort.Strings(above50Games)
	sort.Strings(above60Games)

	fmt.Println("--------------")
	fmt.Println("   ALL GAMES  ")
	fmt.Println("--------------")
	all := strings.Join(allGames, " ")

	fmt.Printf("(greengoo71) 422 : %s\n", strings.Join(above60Games, " "))
	fmt.Printf("(greengoo71) 423 : %s\n", all)
	fmt.Printf("(greengoo71) 424 : %s\n", all)
	fmt.Printf("(greengoo71) 425 : %s\n", strings.Join(above35Games, " "))
	fmt.Printf("(greengoo71) 426 : %s\n", strings.Join(above50Games, " "))
	fmt.Printf("(greengoo71) 427 : %s\n", all)

}

func fetchIfMissing(games map[int]string) {
	fileTemplate := "/home/mapi/work/20171103_go/src/mathandel/gry/%d.xml"
	for id, _ := range games {
		path := fmt.Sprintf(fileTemplate, id)
		if _, err := os.Stat(path); err != nil {
			r := fetchUrl(fmt.Sprintf(gameTemplate, id))
			ioutil.WriteFile(path, r, 0644)
			fmt.Printf("file %d created!\n", id)
		} else {
			fmt.Printf("file %d exists!\n", id)
		}
	}
}

func worker(workerId int, jobs <-chan int, results chan<- GameMetadata) {
	for id := range jobs {
		meta := fetchGameMetadata(id)
		results <- meta
	}
}

func submitAllJobs(jobs chan<- int, gameNames map[int]string) {
	fmt.Printf("there are %d tasks to finish\n", len(gameNames))
	for id, _ := range gameNames {
		jobs <- id
	}
	close(jobs)
}

func fetchGameMetadata(id int) GameMetadata {
	r := fetchUrl(fmt.Sprintf(gameFileTemplate, id))
	var boardgames Boardgames
	xml.Unmarshal(r, &boardgames)
	game := boardgames.Boardgame
	priceUSD, priceEUR := countAveragePrice(game.Listings)

	metadata := GameMetadata{
		id,
		game.Year,
		len(game.Awards),
		game.Mechanics,
		game.Categories,
		game.Rating,
		priceUSD,
		priceEUR}

	return metadata
}

func fetchUrl(url string) []byte {
	if strings.HasPrefix(url, "http") {
		r, err := http.Get(url)
		handleErr(err)
		for r.StatusCode != 200 {
			r.Body.Close()
			fmt.Printf("retry %d\n", r.StatusCode)
			r, err = http.Get(url)
			handleErr(err)
		}
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		handleErr(err)
		return b
	}
	b, err := ioutil.ReadFile(url)
	handleErr(err)
	return b
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func countAveragePrice(listings []Listing) (float64, float64) {
	rPrice, err := regexp.Compile("<price currency='(.+)'>([0-9.]+)</price>")
	handleErr(err)
	rCondition, err := regexp.Compile("<condition>(.+)</condition>")
	handleErr(err)
	m := make(map[string]map[string][]float64)
	for _, l := range listings {
		match := rPrice.FindStringSubmatch(l.Body)
		currency := match[1]
		if currency == "USD" || currency == "EUR" {
			price, err := strconv.ParseFloat(match[2], 64)
			handleErr(err)
			match = rCondition.FindStringSubmatch(l.Body)
			condition := match[1]
			if condition == "new" || condition == "likenew" {
				if m[currency] == nil {
					m[currency] = make(map[string][]float64)
				}
				m[currency][condition] = append(m[currency][condition], price)
			}
		}
	}

	ret := make(map[string]float64)

	for cur, mm := range m {
		for con, prices := range mm {
			var sum float64
			for _, p := range prices {
				sum += p
			}
			ret[fmt.Sprintf("%s,%s", cur, con)] = sum / float64(len(prices))
		}
	}

	var priceUSD float64
	if val, ok := ret["USD,new"]; ok {
		priceUSD = val
	} else if val, ok := ret["USD,likenew"]; ok {
		priceUSD = val
	}

	var priceEUR float64
	if val, ok := ret["EUR,new"]; ok {
		priceEUR = val
	} else if val, ok := ret["EUR,likenew"]; ok {
		priceEUR = val
	}

	return priceUSD, priceEUR
}
