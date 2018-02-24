package main

import (
	"fmt"
	"strings"
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
