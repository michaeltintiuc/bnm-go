package main

import (
	"testing"
)

func Test_fetchURL(t *testing.T) {
	xml, err := fetchURL(buildURL())

	if err != nil {
		t.Error(err)
		return
	}

	if len(xml) == 0 {
		t.Error("Empty response")
	}
}

func Test_parseXML(t *testing.T) {
	const xml = `
<ValCurs Date="02.01.2006" name="Official exchange rate">
	<Valute ID="47">
		<NumCode>978</NumCode>
		<CharCode>EUR</CharCode>
		<Nominal>1</Nominal>
		<Name>Euro</Name>
		<Value>15.1950</Value>
	</Valute>
	<Valute ID="44">
		<NumCode>840</NumCode>
		<CharCode>USD</CharCode>
		<Nominal>1</Nominal>
		<Name>US Dollar</Name>
		<Value>12.8320</Value>
	</Valute>
</ValCurs>`

	r, err := parseXML([]byte(xml))

	if err != nil {
		t.Error(err)
		return
	}

	if len(r.Rates) != 2 {
		t.Error("Expected 2 rates")
		return
	}

	for _, r := range r.Rates {
		if r.CharCode == "" {
			t.Error("Empty value")
			return
		}
		if r.Name == "" {
			t.Error("Empty value")
			return
		}
		if r.NumCode == 0 {
			t.Error("Empty value")
			return
		}
		if r.Value == 0 {
			t.Error("Empty value")
			return
		}
	}
}
