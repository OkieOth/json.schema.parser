package jsonschemaparser_test

import (
	"fmt"
	"os"
	"testing"

	o "github.com/okieoth/goptional/v3"
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
		{"alreadyProper", "AlreadyProper"},
		{"snake_case_test", "SnakeCaseTest"},
		{"kebab-case-test", "KebabCaseTest"},
		{" mixed Case string ", "MixedCaseString"},
		{"!@#$$%^", ""},
		{"goLang", "GoLang"},
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

func TestTopLevelInt(t *testing.T) {
	tests := []struct {
		input    string
		expected types.IntegerType
	}{
		{
			input: "_resources/tests/int_1.json",
			expected: types.IntegerType{
				Name:             o.NewOptionalValue("IAmAnInt"),
				Default:          o.NewOptionalValue(10),
				Minimum:          o.NewOptionalValue(-1),
				ExclusiveMaximum: o.NewOptionalValue(20),
				MultipleOf:       o.NewOptionalValue(2),
			},
		},
		{
			input: "_resources/tests/int_2.json",
			expected: types.IntegerType{
				Name:             o.NewOptionalValue("IAmAnotherInt"),
				Format:           o.NewOptionalValue("uint32"),
				Default:          o.NewOptionalValue(10),
				Maximum:          o.NewOptionalValue(100),
				ExclusiveMinimum: o.NewOptionalValue(-20),
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
			x, ok := v.(types.IntegerType)
			require.True(t, ok)
			require.Equal(t, test.expected, x)
		}
	}
}

func TestTopLevelNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected types.NumberType
	}{
		{
			input: "_resources/tests/number_1.json",
			expected: types.NumberType{
				Name:             o.NewOptionalValue("IAmANumber"),
				Default:          o.NewOptionalValue(10.1),
				Minimum:          o.NewOptionalValue(-1.1),
				ExclusiveMaximum: o.NewOptionalValue(20.5),
			},
		},
		{
			input: "_resources/tests/number_2.json",
			expected: types.NumberType{
				Name:             o.NewOptionalValue("IAmAnotherNumber"),
				Format:           o.NewOptionalValue("float32"),
				Default:          o.NewOptionalValue(10.5),
				Maximum:          o.NewOptionalValue(100.1),
				ExclusiveMinimum: o.NewOptionalValue(-20.0),
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
			x, ok := v.(types.NumberType)
			require.True(t, ok)
			require.Equal(t, test.expected, x)
		}
	}
}

func TestTopLevelBool(t *testing.T) {
	tests := []struct {
		input    string
		expected types.BoolType
	}{
		{
			input: "_resources/tests/bool_1.json",
			expected: types.BoolType{
				Name:    o.NewOptionalValue("IAmABool"),
				Default: o.NewOptionalValue(true),
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
			x, ok := v.(types.BoolType)
			require.True(t, ok)
			require.Equal(t, test.expected, x)
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

func TestTopArrayType(t *testing.T) {
	tests := []struct {
		input    string
		expected types.ArrayType
	}{
		{
			input: "_resources/tests/array_1.json",
			expected: types.ArrayType{
				Name: o.NewOptionalValue("IAmAnArray"),
				ValueType: types.IntegerType{
					Default:          o.NewOptionalValue(10),
					Minimum:          o.NewOptionalValue(-1),
					ExclusiveMaximum: o.NewOptionalValue(20),
					MultipleOf:       o.NewOptionalValue(2),
				},
			},
		},
		{
			input: "_resources/tests/array_2.json",
			expected: types.ArrayType{
				Name:     o.NewOptionalValue("IAmAnotherArray"),
				MinItems: o.NewOptionalValue(1),
				MaxItems: o.NewOptionalValue(10),
				ValueType: types.UUIDType{
					Default: o.NewOptionalValue("7aca7f3e-6cbc-11f0-a324-037c2bd487d4"),
				},
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
			x, ok := v.(types.ArrayType)
			require.True(t, ok)
			require.Equal(t, test.expected, x)
		}
	}
}

func TestTopMapType(t *testing.T) {
	tests := []struct {
		input    string
		expected types.MapType
	}{
		{
			input: "_resources/tests/map_1.json",
			expected: types.MapType{
				Name:     "IAmAMap",
				TopLevel: true,
				ValueType: types.IntegerType{
					Default:          o.NewOptionalValue(10),
					Minimum:          o.NewOptionalValue(-1),
					ExclusiveMaximum: o.NewOptionalValue(20),
					MultipleOf:       o.NewOptionalValue(2),
				},
			},
		},
		{
			input: "_resources/tests/map_2.json",
			expected: types.MapType{
				Name:     "IAmMap2",
				TopLevel: true,
				ValueType: types.ArrayType{
					ValueType: types.IntegerType{
						Default:          o.NewOptionalValue(10),
						Minimum:          o.NewOptionalValue(-1),
						ExclusiveMaximum: o.NewOptionalValue(20),
						MultipleOf:       o.NewOptionalValue(2),
					},
				},
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
			x, ok := v.(types.MapType)
			require.True(t, ok)
			require.Equal(t, test.expected, x)
		}
	}
}
