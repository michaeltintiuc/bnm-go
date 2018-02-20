package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const domain string = "http://bnm.md/"
const endpoint string = "/official_exchange_rates?get_xml=1&date="

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
	return domain + lang + endpoint + time.Format("02.01.2006")
}

func main() {
	var rates Rates
	res, _ := http.Get(buildURL("en", time.Now()))
	bytes, _ := ioutil.ReadAll(res.Body)
	xml.Unmarshal(bytes, &rates)
	fmt.Println(rates.Map())
}
