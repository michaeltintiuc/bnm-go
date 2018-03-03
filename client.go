package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func buildURL() string {
	return domain + lang + endpoint + date
}

// getXML
// Fetch data from a cache file or
// Send new request and cache response
func getXML() ([]byte, error) {
	tmp := fmt.Sprintf("%s/%s-%s-%s", os.TempDir(), "bnm-go", lang, date)

	if _, err := os.Stat(tmp); os.IsNotExist(err) || fresh {
		if verbose {
			if fresh {
				fmt.Printf(">Skipping reading cache file %s\n", tmp)
			} else {
				fmt.Printf(">Cache file %s doesn't exist\n", tmp)
			}
		}

		xml, err := fetchURL(buildURL())

		if err != nil {
			fmt.Printf(">Failed to fetch data\n>%s\n", err)
		} else {
			wg.Add(1)
			go cacheXML(tmp, xml)
			wg.Wait()
		}

		return xml, err
	}

	if verbose {
		fmt.Printf(">Reading from cache file %s\n", tmp)
	}

	return ioutil.ReadFile(tmp)
}

func cacheXML(path string, xml []byte) error {
	mu.Lock()
	err := ioutil.WriteFile(path, xml, 0664)

	if err != nil {
		fmt.Printf(">Failed to write cache to %s\n>%s\n", path, err)
	} else if verbose {
		fmt.Printf(">Created cache file %s\n", path)
	}

	defer wg.Done()
	defer mu.Unlock()

	return err
}

func fetchURL(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode != 200 {
		return []byte{}, errors.New("Unexpected respose " + res.Status)
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func parseXML(bytes []byte) (Rates, error) {
	var rates Rates
	err := xml.Unmarshal(bytes, &rates)

	if err != nil {
		return Rates{}, err
	}

	return rates, err
}
