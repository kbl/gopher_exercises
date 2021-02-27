package my_xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func ensureExists(archiveDirectory string) {
	if _, err := os.Stat(archiveDirectory); err == nil {
		return // dir exists
	}
	if err := os.MkdirAll(archiveDirectory, os.ModeDir|0777); err != nil {
		log.Fatal(err)
	}
}

func ArchiveTo(archiveDirectory string) {
	ensureExists(archiveDirectory)
	newestId := getNewestId()
	missingIds := findMissingComics(archiveDirectory, newestId)

	for _, comicId := range missingIds {
		comic := DownloadIssue(comicId)
		saveInArchive(archiveDirectory, comic)
	}
}

func saveInArchive(archiveDirectory string, comic *Comic) {
	filePath := path.Join(archiveDirectory, fmt.Sprintf("%d.json", comic.Num))

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(comic)
}

func getNewestId() int {
	comic := download(indexURL)
	return comic.Num
}

func findMissingComics(archiveDirectory string, newestId int) []int {
	files, err := ioutil.ReadDir(archiveDirectory)
	if err != nil {
		log.Fatal(err)
	}

	presentFiles := make(map[int]bool)
	for _, f := range files {
		extIndex := len(f.Name()) - len(path.Ext(f.Name()))
		comicId, err := strconv.Atoi(f.Name()[:extIndex])
		if err != nil {
			log.Fatal(err)
		}
		presentFiles[comicId] = true
	}

	var missingComicIds []int

	for comicId := 1; comicId <= newestId; comicId++ {
		if _, exists := presentFiles[comicId]; !exists {
			missingComicIds = append(missingComicIds, comicId)
		}
	}
	return missingComicIds
}

func DownloadIssue(id int) *Comic {
	url := fmt.Sprintf(comicURLTemplate, id)
	return download(url)
}

func download(url string) *Comic {
	comic := new(Comic)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(comic); err != nil {
		log.Fatal(err)
	}

	return comic
}
