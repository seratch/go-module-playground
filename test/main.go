package main

import (
	"fmt"
	hello2 "github.com/seratch/go-module-playground/v2"
	hello3 "github.com/seratch/go-module-playground/v3"
)

func main() {
	fmt.Println(hello3.Version())
	fmt.Println(hello3.Hello())
	fmt.Println(hello2.Version())
}