package main

const (
	domain     string = "http://bnm.md/"
	endpoint   string = "/official_exchange_rates?get_xml=1&date="
	dateFormat string = "02.01.2006"
)

func main() {
	if err := parseFlags(); err != nil || help {
		return
	}

	if err := getRates(false); err != nil {
		return
	}

	if compare {
		if err := getPastRates(); err != nil {
			return
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
