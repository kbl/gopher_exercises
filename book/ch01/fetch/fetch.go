package fetch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("fetch: %v\n", err), err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("fetch: reading %s: %v", url, err), err
	}

	if resp.StatusCode != 200 {
		return fmt.Sprintf("%s", b), errors.New(fmt.Sprintf("Illegal response status %d.", resp.StatusCode))
	}
	return fmt.Sprintf("%s", b), nil
}
