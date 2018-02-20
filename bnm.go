package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const domain string = "http://bnm.md/"
const endpoint string = "/official_exchange_rates?get_xml=1&date="
const dateFormat string = "02.01.2006"

var cache = true

// Rates parent node
type Rates struct {
	Rates []Rate `xml:"Valute"`
}

// Map convert rates slices to a map
func (r Rates) Map() map[string]Rate {
	ratesMap := make(map[string]Rate)
	for i := range r.Rates {
		ratesMap[r.Rates[i].CharCode] = r.Rates[i]
	}
	return ratesMap
}

// Rate child node
type Rate struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float32 `xml:"Value"`
}

func buildURL(lang, date string) string {
	lang = strings.ToLower(lang)
	validateLang(lang)
	return domain + lang + endpoint + date
}

func validateLang(lang string) {
	switch lang {
	case "en", "ru", "ro", "md":
		return
	}
	panic("Invalid language. Supported languages: en, ru, ro, md")
}

// getXML
// Fetch data from a cache file or
// Send new request and cache response
func getXML() ([]byte, error) {
	date := time.Now().Format(dateFormat)
	tmp := os.TempDir() + "/bnm-go-" + date

	if _, err := os.Stat(tmp); os.IsNotExist(err) {
		fmt.Printf(">Cache file %s doesn't exist\n", tmp)
		xml, err := fetchURL(buildURL("en", date))

		if err != nil {
			fmt.Println(">Failed to fetch data", err)
			return []byte{}, err
		}

		err = ioutil.WriteFile(tmp, xml, 0664)
		if err != nil {
			fmt.Printf(">Failed to write to %s\n>%s\n", tmp, err)
		}
		return xml, err
	}

	fmt.Printf(">Reading from %s\n", tmp)
	return ioutil.ReadFile(tmp)
}

func fetchURL(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return []byte{}, err
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

func main() {
	xml, _ := getXML()
	rates, _ := parseXML(xml)

	fmt.Println(rates.Map())
}
