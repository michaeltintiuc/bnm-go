package main

const (
	domain     string = "http://bnm.md/"
	endpoint   string = "/official_exchange_rates?get_xml=1&date="
	dateFormat string = "02.01.2006"
)

func main() {
	parseFlags()
	getRates(false)

	if compare {
		getPastRates()
	}

	printCurrentRates()

	if buy > 0 {
		printBuyRates()
	}

	if sell > 0 {
		printSellRates()
	}
}
