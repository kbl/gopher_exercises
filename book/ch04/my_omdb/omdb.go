package omdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// posters api is only available for patrons
const searchURLTemplate = "http://omdbapi.com/?apikey=%s&&t=%s"

type movieResponse struct {
	Poster string
}

func DownloadPoster(apiToken, movieTitle string) []byte {
	url := fmt.Sprintf(searchURLTemplate, url.QueryEscape(apiToken), url.QueryEscape(movieTitle))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var posterPath movieResponse
	if err := json.NewDecoder(resp.Body).Decode(&posterPath); err != nil {
		log.Fatal(err)
	}
	if posterPath.Poster == "" {
		log.Fatalf("Movie with title %q is missing!", movieTitle)
	}

	posterResponse, err := http.Get(posterPath.Poster)
	if err != nil {
		log.Fatal(err)
	}
	defer posterResponse.Body.Close()
	var posterBuffer bytes.Buffer
	posterBuffer.ReadFrom(posterResponse.Body)

	return posterBuffer.Bytes()
}
