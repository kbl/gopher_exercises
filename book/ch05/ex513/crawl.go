package ex513

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/links"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Crawl(aUrl string) {
	crawl(aUrl, []string{})
}

func crawl(aUrl string, alreadyVisited []string) []string {
	parsedUrl := parseUrl(aUrl)
	aUrl = fmt.Sprintf("%s://%s%s", parsedUrl.Scheme, parsedUrl.Host, parsedUrl.Path)

	if filepath.Ext(parsedUrl.Path) == "" && !strings.HasSuffix(parsedUrl.Path, "/") {
		aUrl = fmt.Sprintf("%s/", aUrl)
	}

	if wasVisited(aUrl, alreadyVisited) {
		return alreadyVisited
	}

	alreadyVisited = append(alreadyVisited, aUrl)

	list, err := links.Extract(aUrl)
	if err != nil {
		log.Fatal(err)
	}

	downloadPage(parsedUrl, aUrl)

	for _, link := range list {
		if parseUrl(link).Host != parsedUrl.Host {
			log.Printf("Skipping %s.\n", link)
			continue
		}
		alreadyVisited = crawl(link, alreadyVisited)
	}

	return alreadyVisited
}

func downloadPage(parsedUrl *url.URL, aUrl string) {
	fmt.Printf("Downloading url %s.\n", aUrl)
	resp, err := http.Get(aUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("getting %s: %s", aUrl, resp.Status)
		os.Exit(1)
	}

	dirName, fileName := filepath.Split(fmt.Sprintf("%s%s", parsedUrl.Host, parsedUrl.Path))
	if dirName == "" {
		dirName = fileName
		fileName = "index.html"
	}
	if fileName == "" {
		fileName = "index.html"
	}
	filePath := filepath.Join(dirName, fileName)
	fmt.Println(dirName, fileName, filePath)

	if _, err = os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func wasVisited(aUrl string, alreadyVisited []string) bool {
	for _, visited := range alreadyVisited {
		if visited == aUrl {
			return true
		}
	}
	return false
}

func parseUrl(aUrl string) *url.URL {
	parsedUrl, err := url.Parse(aUrl)
	if err != nil {
		log.Fatal(err)
	}
	return parsedUrl
}
