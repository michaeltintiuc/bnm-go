package main

import (
	"time"
)

var (
	ratesUsed = []Rate{}
	ratesPast = []Rate{}
)

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

func getRates(past bool) {
	xml, _ := getXML()
	rates, _ := parseXML(xml)

	for _, r := range rates.Rates {
		if currencies.Contains(r.CharCode) {
			if past {
				ratesPast = append(ratesPast, r)
			} else {
				ratesUsed = append(ratesUsed, r)
			}
		}
	}
}

func getPastRates() {
	dateTime, _ := time.Parse(dateFormat, date)
	dateBak := date
	date = dateTime.AddDate(0, 0, -1).Format(dateFormat)
	getRates(true)
	date = dateBak
}
