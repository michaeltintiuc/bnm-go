package main

import (
	"net/http"
	"testing"
)

func TestPingURL(t *testing.T) {
	res, err := http.Get(buildURL())

	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected 200 OK response, received %s", res.Status)
	}
}

func Test_fetchURL(t *testing.T) {
	xml, err := fetchURL(buildURL())

	if err != nil {
		t.Fatal(err)
	}

	if xml == nil || len(xml) == 0 {
		t.Fatal("Empty response")
	}
}

func Test_fetchURL_404(t *testing.T) {
	xml, err := fetchURL(domain + "/foobar")

	if err == nil {
		t.Fatal("Expected error")
	}

	if len(xml) > 0 {
		t.Fatal("Expected empy result")
	}
}

func Test_fetchURL_faulty(t *testing.T) {
	_, err := fetchURL("foobar")

	if err == nil {
		t.Fatal("Expected error")
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
		t.Fatal(err)
	}

	if len(r.Rates) != 2 {
		t.Fatalf("Expected 2 rates, received %d", len(r.Rates))
	}

	for _, r := range r.Rates {
		if r.CharCode == "" {
			t.Fatal("Empty CharCode value")
		}
		if r.Name == "" {
			t.Fatal("Empty Name value")
		}
		if r.NumCode == 0 {
			t.Fatal("Empty NumCode value")
		}
		if r.Value == 0 {
			t.Fatal("Empty Value value")
		}
	}
}

func Test_parseXML_empty(t *testing.T) {
	r, _ := parseXML([]byte{})

	if r.Rates != nil {
		t.Fatal("Expected error")
	}
}

func Test_cacheXML_faulty(t *testing.T) {
	wg.Add(1)
	if err := cacheXML("/foobar", []byte{}); err == nil {
		t.Fatal("Expected error")
	}
	wg.Wait()
}
