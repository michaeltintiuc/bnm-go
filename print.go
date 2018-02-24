package main

import (
	"fmt"
)

const (
	clrR string = "\x1b[31;1m"
	clrG string = "\x1b[32;1m"
	clrB string = "\x1b[34;1m"
	clrN string = "\x1b[0m"
)

func printCurrentRates() {
	fmt.Printf("Rates for %s:\n", date)

	for _, r := range ratesUsed {
		if compare {
			clr := clrB

			for _, p := range ratesPast {
				if r.CharCode == p.CharCode {
					if r.Value < p.Value {
						clr = clrR
					} else if r.Value > p.Value {
						clr = clrG
					}
					break
				}
			}
			fmt.Printf("%s%s %f (%s)%s\n", clr, r.CharCode, r.Value, r.Name, clrN)
		} else {
			fmt.Printf("%s %f (%s)\n", r.CharCode, r.Value, r.Name)
		}
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
