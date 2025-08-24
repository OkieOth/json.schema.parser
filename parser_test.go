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
	result, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	require.Equal(t, 4, len(result.ComplexTypes), "wrong number of elements in types map")
	require.Equal(t, 1, len(result.StringEnums), "wrong number of elements in types string enum")
	require.Equal(t, 0, len(result.ArrayTypes))
	require.Equal(t, 0, len(result.MapTypes))
	require.Equal(t, 0, len(result.IntEnums))
	require.Equal(t, 0, len(result.IntegerTypes))
	require.Equal(t, 0, len(result.NumberTypes))
	require.Equal(t, 0, len(result.StringTypes))
	require.Equal(t, 0, len(result.UUIDTypes))
	require.Equal(t, 0, len(result.DateTypes))
	require.Equal(t, 0, len(result.DateTimeTypes))
	require.Equal(t, 0, len(result.TimeTypes))
	require.Equal(t, 0, len(result.DurationTypes))
	require.Equal(t, 0, len(result.BoolTypes))
	require.Equal(t, 0, len(result.BinaryTypes))
	require.Equal(t, 0, len(result.ObjectTypes))

	personName, exist := result.ComplexTypes["PersonName"]
	require.True(t, exist, "PersonName doesn't exist")
	fmt.Println(personName)

	personContact, exist := result.ComplexTypes["PersonContact"]
	require.True(t, exist, "PersonContact doesn't exist")
	fmt.Println(personContact)

	personContactAddress, exist := result.ComplexTypes["PersonContactAddress"]
	require.True(t, exist, "PersonContactAddress doesn't exist")
	fmt.Println(personContactAddress)

	personRoles, exist := result.StringEnums["PersonRolesItems"]
	require.True(t, exist, "PersonRolesItems doesn't exist")
	fmt.Println(personRoles)

	_, exist = result.ComplexTypes["Person"]
	require.True(t, exist, "Person doesn't exist")
}

func TestParseBytesWithRefs(t *testing.T) {
	bytes, err := os.ReadFile("_resources/tests/test_schema_with_refs.json")
	require.Nil(t, err)
	result, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	require.Equal(t, 4, len(result.ComplexTypes), "wrong number of elements in types map")
	require.Equal(t, 1, len(result.StringEnums), "wrong number of elements in types string enum")
	require.Equal(t, 0, len(result.ArrayTypes))
	require.Equal(t, 0, len(result.MapTypes))
	require.Equal(t, 0, len(result.IntEnums))
	require.Equal(t, 0, len(result.IntegerTypes))
	require.Equal(t, 0, len(result.NumberTypes))
	require.Equal(t, 0, len(result.StringTypes))
	require.Equal(t, 0, len(result.UUIDTypes))
	require.Equal(t, 0, len(result.DateTypes))
	require.Equal(t, 0, len(result.DateTimeTypes))
	require.Equal(t, 0, len(result.TimeTypes))
	require.Equal(t, 0, len(result.DurationTypes))
	require.Equal(t, 0, len(result.BoolTypes))
	require.Equal(t, 0, len(result.BinaryTypes))
	require.Equal(t, 0, len(result.ObjectTypes))

	personName, exist := result.ComplexTypes["Name"]
	require.True(t, exist, "PersonName doesn't exist")
	fmt.Println(personName)

	personContact, exist := result.ComplexTypes["PersonContact"]
	require.True(t, exist, "PersonContact doesn't exist")
	fmt.Println(personContact)

	personContactAddress, exist := result.ComplexTypes["PersonContactAddress"]
	require.True(t, exist, "PersonContactAddress doesn't exist")
	fmt.Println(personContactAddress)

	personRoles, exist := result.StringEnums["PersonRolesItems"]
	require.True(t, exist, "PersonRolesItems doesn't exist")
	fmt.Println(personRoles)

	person, exist := result.ComplexTypes["Person"]
	require.True(t, exist, "Person doesn't exist")
	found := false
	for _, p := range person.Properties {
		if p.Name == "name" {
			found = true
			nameType, ok := p.ValueType.(types.ComplexType)
			require.True(t, ok, "name property of Person isn't of ComplexType")
			require.Equal(t, "Name", nameType.Name, "name property of Person doesn't point to a type with name 'Name'")
		}
	}
	require.True(t, found, "didn't find name property in Person type")
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
		require.Len(t, m.IntegerTypes, 1, "wrong number of returned types")
		for _, v := range m.IntegerTypes {
			require.Equal(t, test.expected, v)
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
		require.Len(t, m.NumberTypes, 1, "wrong number of returned types")
		for _, v := range m.NumberTypes {
			require.Equal(t, test.expected, v)
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
		require.Len(t, m.BoolTypes, 1, "wrong number of returned types")
		for _, v := range m.BoolTypes {
			require.Equal(t, test.expected, v)
		}
	}
}

func TestTopLevelEnum(t *testing.T) {
	tests := []struct {
		input     string
		checkFunc func(v types.ParseResult) bool
	}{
		{
			input: "_resources/tests/enum_str_1.json",
			checkFunc: func(v types.ParseResult) bool {
				require.Equal(t, 1, len(v.StringEnums))
				for _, x := range v.StringEnums {
					assert.Equal(t, []string{"red", "green", "blue"}, x.Values)
				}
				return true
			},
		},
		{
			input: "_resources/tests/enum_str_2.json",
			checkFunc: func(v types.ParseResult) bool {
				require.Equal(t, 1, len(v.StringEnums))
				for _, x := range v.StringEnums {
					assert.Equal(t, []string{"red", "green", "blue"}, x.Values)
				}
				return true
			},
		},
		{
			input: "_resources/tests/enum_int_1.json",
			checkFunc: func(v types.ParseResult) bool {
				require.Equal(t, 1, len(v.IntEnums))
				for _, x := range v.IntEnums {
					assert.Equal(t, []int{13, 700, 42}, x.Values)
				}
				return true
			},
		},
		{
			input: "_resources/tests/enum_int_2.json",
			checkFunc: func(v types.ParseResult) bool {
				require.Equal(t, 1, len(v.IntEnums))
				for _, x := range v.IntEnums {
					assert.Equal(t, []int{13, 700, 42}, x.Values)
				}
				return true
			},
		},
	}
	for _, test := range tests {
		bytes, err := os.ReadFile(test.input)
		require.Nil(t, err)
		m, err := p.ParseBytes(bytes)
		require.Nil(t, err)
		require.True(t, test.checkFunc(m))
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
		require.Len(t, m.ArrayTypes, 1, "wrong number of returned types")
		for _, v := range m.ArrayTypes {
			require.Equal(t, test.expected, v)
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
		require.Len(t, m.MapTypes, 1, "wrong number of returned types")
		for _, v := range m.MapTypes {
			require.Equal(t, test.expected, v)
		}
	}
}
