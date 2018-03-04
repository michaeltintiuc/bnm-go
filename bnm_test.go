package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	compare, buy, sell = true, 1, 1
	main()
	compare = false

	help = true
	main()
	help = false

	date = "01.01.1970"
	main()
	date = "02.01.2006"
}
