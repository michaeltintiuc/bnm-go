package main

import (
	"testing"
)

func setDefaults() {
	lang = "en"
	date = "02.01.2006"
	buy = 0
	sell = 0
	verbose = false
	help = false
	fresh = false
	compare = false
}

func Test_parseFlags(t *testing.T) {
	setDefaults()
	err := parseFlags()

	if err != nil {
		t.Fatal(err)
	}
}

func Test_langFlag(t *testing.T) {
	setDefaults()
	lang = "qwe"
	err := parseFlags()

	if err == nil {
		t.Fatal("Expected invalid language error")
	}
}

func Test_dateFlag(t *testing.T) {
	setDefaults()
	date = "yd"
	err := parseFlags()

	if err != nil {
		t.Fatal(err)
	}
}

func Test_buyFlag(t *testing.T) {
	setDefaults()
	buy = -1
	sell = -1
	err := parseFlags()

	if err == nil {
		t.Fatal("Expected invalid number error")
	}
}

func Test_helpFlag(t *testing.T) {
	setDefaults()
	help = true
	err := parseFlags()

	if err != nil {
		t.Fatal(err)
	}
}
