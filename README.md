# bnm-go
a CLI tool for fetching official currency rates from the [National Bank of Moldova](https://bnm.md/)

## Installation

`go get github.com/michaeltintiuc/bnm-go`

![bnm-go preview](/preview.gif?raw=true "bnm-go preview")

## Usage

`bnm-go -h`

```
  -buy float
    	Calculate amount of MDL for each -c (currencies) bought
  -c value
    	Comma separated list of currencies to display (default [USD])
  -d string
    	Date format: {dd.mm.yyy} or {yesterday|yday|yd|yda} (default "24.02.2018")
  -f	Skip reading cache and fetch fresh data
  -h	Print usage
  -l string
    	Language: {en|md|ro|ru} (default "en")
  -sell float
    	Calculate amount of MDL for each -c (currencies) sold
  -v	Display verbose output
  -x	Cross reference rates to the day before -d
```

Produce verbose output

`bnm-go -v`

Fetch fresh data, skipping any cached files

`bnm-go -f`

Fetch rates and set output language to Russian (ru)

`bnm-go -l=ru`

Fetch USD and EUR rates for _today_

`bnm-go -c=usd,eur`

Alternate way of setting currencies

`bnm-go -c=eur -c=cad -c=nzd`

Fetch USD and EUR rates for yesterday

`bnm-go -c=usd,eur -d=yd`

Fetch USD and EUR rates for January 31st 20018

`bnm-go -c=usd,eur -d=31.01.2018`

Fetch USD and EUR rates and compare with values from yesterday or day prior to `-d`. The output will be display red, green and blue colors for drop, rise and unchanged respectively

`bnm-go -c=usd,eur -x`

Calculate amount of MDL received for selling 100 USD and EUR

`bnm-go -c=usd,eur --sell=100`

Calculate amount of USD and EUR received for 100 MDL

`bnm-go -c=usd,eur --buy=100`


## TODO:
- [X] Caching
- [X] Argument parsing (date, currency)
- [X] Currency conversion
- [X] Cross reference rates from prior date for drop/rise (red/green) display
- [ ] Tests
- [ ] Docs
