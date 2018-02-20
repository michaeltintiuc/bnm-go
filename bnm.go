package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const domain string = "http://bnm.md/"
const endpoint string = "/official_exchange_rates?get_xml=1&date="
const dateFormat string = "02.01.2006"

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

func buildURL(lang string, time time.Time) string {
	lang = strings.ToLower(lang)
	validateLang(lang)
	return domain + lang + endpoint + time.Format(dateFormat)
}

func validateLang(lang string) {
	switch lang {
	case "en", "ru", "ro", "md":
		return
	}
	panic("Invalid language. Supported languages: en, ru, ro, md")
}

func fetchURL(url string) ([]byte, error) {
	res, err := http.Get(buildURL("en", time.Now()))

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
	res, _ := fetchURL(buildURL("en", time.Now()))
	rates, _ := parseXML(res)

	fmt.Println(rates.Map())
}
