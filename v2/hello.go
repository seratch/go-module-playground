package hello

import "rsc.io/quote"

func Version() string {
	return "2.0.0"
}

func Hello() string {
	return quote.Hello()
}
