package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var (
	date, lang     string
	buy, sell      float64
	verbose, fresh bool
	help, compare  bool
	currencies     currencySlice
)

func validateFlags() error {
	// Language validation
	lang = strings.ToLower(lang)
	switch lang {
	case "en", "ru", "ro", "md":
	default:
		return errors.New("Invalid language \"" + lang + "\" provided")
	}

	// Date validation
	switch strings.ToLower(date) {
	case "yesterday", "yday", "yd", "yda":
		date = time.Now().AddDate(0, 0, -1).Format(dateFormat)
	}

	// Buy,Sell validation
	if buy < 0 || sell < 0 {
		return errors.New("Negative numbers are not supported")
	}

	// Set default currency
	if len(currencies) == 0 {
		currencies = append(currencies, "USD")
	}

	return nil
}

func parseFlags() error {
	flag.Parse()
	err := validateFlags()

	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
	}

	if help {
		flag.PrintDefaults()
	}

	return err
}

func init() {
	flag.StringVar(&date, "d", time.Now().Format(dateFormat), "Date format: {dd.mm.yyyy} or {yesterday|yday|yd|yda}")
	flag.StringVar(&lang, "l", "en", "Language: {en|md|ro|ru}")
	flag.Float64Var(&buy, "buy", 0, "Calculate amount of MDL for each -c (currencies) bought")
	flag.Float64Var(&sell, "sell", 0, "Calculate amount of MDL for each -c (currencies) sold")
	flag.Var(&currencies, "c", "Comma separated list of currencies to display")
	flag.BoolVar(&verbose, "v", false, "Display verbose output")
	flag.BoolVar(&fresh, "f", false, "Skip reading cache and fetch fresh data")
	flag.BoolVar(&help, "h", false, "Print usage")
	flag.BoolVar(&compare, "x", false, "Cross reference rates to the day before -d")
}
