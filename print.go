package main

import (
	"fmt"
)

var (
	clrR = "\x1b[31;1m"
	clrG = "\x1b[32;1m"
	clrB = "\x1b[34;1m"
	clrN = "\x1b[0m"
)

func printCurrentRates() {
	fmt.Printf("Rates for %s:\n", date)

	for _, r := range ratesUsed {
		if compare {
			clr := clrB
			arr := "\u2194" // ↔
			d := 0.0

			for _, p := range ratesPast {
				if r.CharCode == p.CharCode {
					d = r.Value - p.Value
					if r.Value < p.Value {
						clr = clrR
						arr = "\u2193" // ↓
					} else if r.Value > p.Value {
						clr = clrG
						arr = "\u2191" // ↑
					}
					break
				}
			}
			fmt.Printf("%s%s %s %.4f (%s) (%.4f)%s\n", clr, arr, r.CharCode, r.Value, r.Name, d, clrN)
		} else {
			fmt.Printf("%s %.4f (%s)\n", r.CharCode, r.Value, r.Name)
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
