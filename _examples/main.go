package main

import (
	"fmt"
	"os"

	p "github.com/okieoth/json.schema.parser"
)

func main() {
	bytes, err := os.ReadFile("test_schema.json")
	if err != nil {
		panic(fmt.Sprintf("can't read input file: %v", err))
	}
	types, err := p.ParseBytes(bytes)
	if err != nil {
		panic(fmt.Sprintf("can't parse bytes: %v", err))
	}
	fmt.Println(types)
}
