package main

import (
	"testing"
)

func Test_currencySlice_Contains(t *testing.T) {
	currencies := currencySlice{"usd", "eur", "cad"}

	if !currencies.Contains("usd") {
		t.Fatal("Expected true, got false")
	}

	if currencies.Contains("qwe") {
		t.Fatal("Expected false, got true")
	}
}

func Test_currencySlice_Set(t *testing.T) {
	currencies := currencySlice{}
	err := currencies.Set("usd,eur,cad,usd,")

	if err != nil {
		t.Fatal(err)
	}
}
