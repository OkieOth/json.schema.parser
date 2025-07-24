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

func TestToProperName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"user-profile_page", "UserProfilePage"},
		{"123 go!", "123Go"},
		{"alreadyProper", "Alreadyproper"},
		{"snake_case_test", "SnakeCaseTest"},
		{"kebab-case-test", "KebabCaseTest"},
		{" mixed CASE string ", "MixedCaseString"},
		{"!@#$$%^", ""},
		{"goLang", "Golang"},
		{"multiple   spaces", "MultipleSpaces"},
		{"with123numbers", "With123numbers"},
	}

	for _, tt := range tests {
		result := p.ToProperName(tt.input)
		if result != tt.expected {
			t.Errorf("ToProperName(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}
