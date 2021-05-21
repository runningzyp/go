package main

import (
	"github.com/runningzyp/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
	println(quote.Hello())
}
