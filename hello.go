package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Version() string {
	return "1.0.2"
}

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
