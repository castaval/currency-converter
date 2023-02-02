package main

import (
	"currency-converter/internal/parser"
	"fmt"
)

func main() {
	data := parser.GetData()
	fmt.Println(data)
}
