package main

import (
	"fmt"
	"os"

	p "github.com/okieoth/jsonschemaparser"
)

func main() {
	fileToUse := "../_resources/tests/test_schema.json"
	bytes, err := os.ReadFile(fileToUse)
	if err != nil {
		panic(fmt.Sprintf("can't read input file: %v", err))
	}
	types, err := p.ParseBytes(bytes)
	if err != nil {
		panic(fmt.Sprintf("can't parse bytes: %v", err))
	}
	fmt.Println(types)
}
