package jsonschemaparser_test

import (
	"fmt"
	"os"
	"testing"

	p "github.com/okieoth/jsonschemaparser"
	"github.com/stretchr/testify/require"
)

func TestParseBytes(t *testing.T) {
	bytes, err := os.ReadFile("_examples/test_schema.json")
	require.Nil(t, err)
	types, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	fmt.Println(types)
}
