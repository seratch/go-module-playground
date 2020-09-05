package main

import (
	"fmt"
	hello2 "github.com/seratch/go-module-playground/v2"
	hello3 "github.com/seratch/go-module-playground/v3"
	hello4 "github.com/seratch/go-module-playground/v4"
)

func main() {
	fmt.Println(hello4.Version())
	fmt.Println(hello3.Version())
	fmt.Println(hello3.Hello())
	fmt.Println(hello2.Version())
}