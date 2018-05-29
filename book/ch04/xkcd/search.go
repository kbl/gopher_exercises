package xkcd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func Search(archiveDirectory, searchTerm string) []*Comic {
	files, err := ioutil.ReadDir(archiveDirectory)
	if err != nil {
		log.Fatal(err)
	}

	var comics []*Comic
	searchTerm = strings.ToLower(searchTerm)

	for _, file := range files {
		var comic Comic
		handle, err := os.Open(path.Join(archiveDirectory, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(handle).Decode(&comic)
		handle.Close()

		transcriptMatches := strings.Index(searchTerm, strings.ToLower(comic.Transcript)) != -1
		titleMatches := strings.Index(searchTerm, strings.ToLower(comic.Title)) != -1

		if transcriptMatches || titleMatches {
			comics = append(comics, &comic)
		}
	}

	return comics
}
