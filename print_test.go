package main

func getRatesUsed() []Rate {
	return []Rate{
		{
			NumCode:  1,
			CharCode: "USD",
			Name:     "US Dollar",
			Value:    16.6802,
		},
		{
			NumCode:  1,
			CharCode: "EUR",
			Name:     "Euro",
			Value:    20.3266,
		},
	}
}

func getRatesPast() []Rate {
	return []Rate{
		{
			NumCode:  1,
			CharCode: "USD",
			Name:     "US Dollar",
			Value:    15.6802,
		},
		{
			NumCode:  1,
			CharCode: "EUR",
			Name:     "Euro",
			Value:    21.3266,
		},
	}
}

func Example_printCurrentRates() {
	ratesUsed = getRatesUsed()
	date = "02.01.2006"
	compare = false
	printCurrentRates()

	// Output:
	// Rates for 02.01.2006:
	// USD 16.6802 (US Dollar)
	// EUR 20.3266 (Euro)
}

func Example_printCurrentRates_compare() {
	ratesUsed = getRatesUsed()
	ratesPast = getRatesPast()
	date = "02.01.2006"
	compare = true
	clrR, clrG, clrB, clrN = "", "", "", ""

	printCurrentRates()

	// Output:
	// Rates for 02.01.2006:
	// USD 16.6802 (US Dollar)
	// EUR 20.3266 (Euro)
}

func Example_printBuyRates() {
	buy = 1
	ratesUsed = getRatesUsed()
	printBuyRates()

	// Output:
	// Buy rates:
	// 1.0000 MDL = 0.0600 USD (US Dollar)
	// 1.0000 MDL = 0.0492 EUR (Euro)
}

func Example_printSellRates() {
	sell = 1
	ratesUsed = getRatesUsed()
	printSellRates()

	// Output:
	// Sell rates:
	// 1.0000 USD (US Dollar) = 16.6802 MDL
	// 1.0000 EUR (Euro) = 20.3266 MDL
}
