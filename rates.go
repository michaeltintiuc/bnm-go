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
	NumCode  int16   `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

func getRates(past bool) error {
	xml, err := getXML()
	if err != nil {
		return err
	}

	rates, err := parseXML(xml)
	if err != nil {
		return err
	}

	for _, r := range rates.Rates {
		if currencies.Contains(r.CharCode) {
			if past {
				ratesPast = append(ratesPast, r)
			} else {
				ratesUsed = append(ratesUsed, r)
			}
		}
	}

	return nil
}

func getPastRates() error {
	dateTime, _ := time.Parse(dateFormat, date)
	dateBak := date
	date = dateTime.AddDate(0, 0, -1).Format(dateFormat)
	err := getRates(true)
	date = dateBak

	return err
}
