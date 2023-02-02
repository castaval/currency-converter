package main

import (
	"currency-converter/internal/currency"
	"currency-converter/internal/parser"
	"fmt"
	"log"
)

func main() {
	data := parser.GetData()
	value, err := currency.Convert(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(value)
}
