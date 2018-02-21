package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	domain     string = "http://bnm.md/"
	endpoint   string = "/official_exchange_rates?get_xml=1&date="
	dateFormat string = "02.01.2006"
)

var (
	date, lang string
)

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

func buildURL() string {
	return domain + lang + endpoint + date
}

// getXML
// Fetch data from a cache file or
// Send new request and cache response
func getXML() ([]byte, error) {
	tmp := fmt.Sprintf("%s/%s-%s-%s", os.TempDir(), "bnm-go", lang, date)

	if _, err := os.Stat(tmp); os.IsNotExist(err) {
		fmt.Printf(">Cache file %s doesn't exist\n", tmp)
		xml, err := fetchURL(buildURL())

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

func validateFlags() {
	lang = strings.ToLower(lang)
	switch lang {
	case "en", "ru", "ro", "md":
	default:
		fmt.Fprintf(os.Stderr, "Invalid language \"%s\" provided\n", lang)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func init() {
	flag.StringVar(&date, "d", time.Now().Format(dateFormat), "Date format: dd.mm.yyy")
	flag.StringVar(&lang, "l", "en", "Language: {en|md|ro|ru}")
}

func main() {
	flag.Parse()
	validateFlags()

	xml, _ := getXML()
	rates, _ := parseXML(xml)

	fmt.Println(rates.Map())
}
