package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	domain     string = "http://bnm.md/"
	endpoint   string = "/official_exchange_rates?get_xml=1&date="
	dateFormat string = "02.01.2006"
)

var (
	wg                   sync.WaitGroup
	mu                   sync.Mutex
	date, lang           string
	buy, sell            float64
	verbose, fresh, help bool
	currencies           = currencySlice{"USD"}
	ratesUsed            = []Rate{}
)

type currencySlice []string

func (c *currencySlice) Contains(needle string) bool {
	for _, value := range *c {
		if needle == value {
			return true
		}
	}
	return false
}

func (c *currencySlice) Set(value string) error {
OUTER:
	for _, currency := range strings.Split(strings.ToUpper(value), ",") {
		if currency == "" {
			continue
		}

		for _, existing := range *c {
			if currency == existing {
				continue OUTER
			}
		}

		*c = append(*c, currency)
	}

	return nil
}

func (c *currencySlice) String() string {
	return fmt.Sprint(*c)
}

// Rates parent node of base XML structure
type Rates struct {
	Rates []Rate `xml:"Valute"`
}

// Rate child node of Rates
type Rate struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

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

	switch strings.ToLower(date) {
	case "yesterday", "yday", "yd", "yda":
		date = time.Now().AddDate(0, 0, -1).Format(dateFormat)
	}

	if buy < 0 || sell < 0 {
		fmt.Fprintf(os.Stderr, "Negative numbers are not supported\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func printCurrentRates() {
	fmt.Printf("Rates for %s:\n", date)
	for _, r := range ratesUsed {
		fmt.Printf("%s %f (%s)\n", r.CharCode, r.Value, r.Name)
	}
}

func printBuyRates() {
	fmt.Println("\nBuy rates:")
	for _, r := range ratesUsed {
		fmt.Printf("%.4f MDL = %.4f %s (%s)\n", buy, buy/r.Value, r.CharCode, r.Name)
	}
}

func printSellRates() {
	fmt.Println("\nSell rates:")
	for _, r := range ratesUsed {
		fmt.Printf("%.4f %s (%s) = %.4f MDL\n", sell, r.CharCode, r.Name, r.Value*sell)
	}
}

func init() {
	flag.StringVar(&date, "d", time.Now().Format(dateFormat), "Date format: {dd.mm.yyy} or {yesterday|yday|yd|yda}")
	flag.StringVar(&lang, "l", "en", "Language: {en|md|ro|ru}")
	flag.Float64Var(&buy, "buy", 0, "Calculate amount of MDL for each -c (currencies) bought")
	flag.Float64Var(&sell, "sell", 0, "Calculate amount of MDL for each -c (currencies) sold")
	flag.Var(&currencies, "c", "Comma separated list of currencies to display")
	flag.BoolVar(&verbose, "v", false, "Display verbose output")
	flag.BoolVar(&fresh, "f", false, "Skip reading cache and fetch fresh data")
	flag.BoolVar(&help, "h", false, "Print usage")
}

func main() {
	flag.Parse()
	validateFlags()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	xml, _ := getXML()
	rates, _ := parseXML(xml)

	for _, r := range rates.Rates {
		if currencies.Contains(r.CharCode) {
			ratesUsed = append(ratesUsed, r)
		}
	}

	printCurrentRates()

	if buy > 0 {
		printBuyRates()
	}

	if sell > 0 {
		printSellRates()
	}
}
