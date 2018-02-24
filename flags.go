package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	date, lang     string
	buy, sell      float64
	verbose, fresh bool
	help, compare  bool
	currencies     = currencySlice{"USD"}
)

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

func parseFlags() {
	flag.Parse()

	validateFlags()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
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
	flag.BoolVar(&compare, "x", false, "Cross reference rates to the day before -d")
}
