package my_xkcd

import "fmt"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const indexURL = "https://xkcd.com/info.0.json"
const comicURLTemplate = "http://xkcd.com/%d/info.0.json"

func (c *Comic) URL() string {
	return fmt.Sprintf(comicURLTemplate, c.Num)
}
