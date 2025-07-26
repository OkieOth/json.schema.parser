package jsonschemaparser_test

import (
	"fmt"
	"os"
	"testing"

	p "github.com/okieoth/jsonschemaparser"
	"github.com/okieoth/jsonschemaparser/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBytes(t *testing.T) {
	bytes, err := os.ReadFile("_resources/tests/test_schema.json")
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

func TestTopLevelEnum(t *testing.T) {
	tests := []struct {
		input     string
		checkFunc func(v any) bool
	}{
		{
			input: "_resources/tests/enum_str_1.json",
			checkFunc: func(v any) bool {
				x, ok := v.(types.StringEnumType)
				if !ok {
					return false
				}
				assert.Equal(t, []string{"red", "green", "blue"}, x.Values)
				return true
			},
		},
		{
			input: "_resources/tests/enum_str_2.json",
			checkFunc: func(v any) bool {
				x, ok := v.(types.StringEnumType)
				if !ok {
					return false
				}
				assert.Equal(t, []string{"red", "green", "blue"}, x.Values)
				return true
			},
		},
		{
			input: "_resources/tests/enum_int_1.json",
			checkFunc: func(v any) bool {
				x, ok := v.(types.IntEnumType)
				if !ok {
					return false
				}
				assert.Equal(t, []int{13, 700, 42}, x.Values)
				return true
			},
		},
		{
			input: "_resources/tests/enum_int_2.json",
			checkFunc: func(v any) bool {
				x, ok := v.(types.IntEnumType)
				if !ok {
					return false
				}
				assert.Equal(t, []int{13, 700, 42}, x.Values)
				return true
			},
		},
	}
	for _, test := range tests {
		bytes, err := os.ReadFile(test.input)
		require.Nil(t, err)
		m, err := p.ParseBytes(bytes)
		require.Nil(t, err)
		require.Len(t, m, 1, "wrong number of returned types")
		for _, v := range m {
			require.True(t, test.checkFunc(v))
		}
	}
}
