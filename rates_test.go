package main

import (
	"testing"
)

func Test_getRates(t *testing.T) {
	err := getRates(false)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_getRates_faulty(t *testing.T) {
	date = "01.01.1970"
	err := getRates(false)

	if err == nil {
		t.Fatal("Expected error")
	}

	date = "02.01.2006"
}

func Test_getPastRates(t *testing.T) {
	err := getPastRates()

	if err != nil {
		t.Fatal(err)
	}
}
