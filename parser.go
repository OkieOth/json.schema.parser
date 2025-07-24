package jsonschemaparser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/okieoth/jsonschemaparser/types"
)

func ParseBytes(input []byte) (map[string]any, error) {
	var parsedSchema map[string]any
	extractedTypes := make(map[string]any, 0)

	if err := json.Unmarshal(input, &parsedSchema); err != nil {
		return extractedTypes, fmt.Errorf("error while unmarshalling schema: %v", err)
	}

	if definitions, ok := parsedSchema["definitions"]; ok {
		definitionsMap, ok := definitions.(map[string]any)
		if !ok {
			return extractedTypes, fmt.Errorf("error while converting to definitions map")
		}
		err := parseTypesFromDefinition(definitionsMap, extractedTypes)
		if err != nil {
			return extractedTypes, fmt.Errorf("error while parsing types in the definitions section: %v", err)
		}
	}

	if _, ok := parsedSchema["properties"]; ok {
		// to level object
		err := parseTopLevelType(parsedSchema, extractedTypes)
		if err != nil {
			return extractedTypes, fmt.Errorf("error while parsing main type: %v", err)
		}
	}

	// TODO - tidy up extractedTypes and remove all redefinitions of basic types (int, number, bool, string)
	return extractedTypes, nil
}

func ToProperName(input string) string {
	// Split the string into fields by any non-letter or digit character
	var result strings.Builder
	capNext := true

	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			capNext = true
			continue
		}
		if capNext {
			result.WriteRune(unicode.ToUpper(r))
			capNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func parseTopLevelType(parsedSchema map[string]any, alreadyExtractedTypes map[string]any) error {
	var typeName string
	if titleEntry, ok := parsedSchema["title"]; ok {
		if t, ok := titleEntry.(string); !ok {
			return fmt.Errorf("title entry of the schema isn't a string")
		} else {
			typeName = ToProperName(t)
		}
	} else {
		currentDate := time.Now().Format("20060102")
		typeName = "UnknownTitle_" + currentDate
	}
	return extractType(typeName, parsedSchema, alreadyExtractedTypes, true)
}

func parseTypesFromDefinition(definitionsMap map[string]any, alreadyExtractedTypes map[string]any) error {
	for typeName, v := range definitionsMap {
		valuesMap, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("entry in definitions map, isn't a map type")
		}
		err := extractType(typeName, valuesMap, alreadyExtractedTypes, true)
		if err != nil {
			return fmt.Errorf("error while extracting type: %v", err)
		}
	}
	return nil
}

func extractType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) error {
	if e, ok := valuesMap["enum"]; ok {
		// found enum entry
		return extractEnumType(name, valuesMap, alreadyExtractedTypes, topLevel, e)
	} else if r, ok := valuesMap["$ref"]; ok {
		// found ref entry
		var refStr string
		if s, ok := r.(string); ok {
			refStr = s
		} else {
			return fmt.Errorf("$ref doesn't point to a string entry, type: %s", name)
		}
		return extractRefType(name, valuesMap, alreadyExtractedTypes, topLevel, refStr)
	} else if t, ok := valuesMap["type"]; ok {
		// found type entry
		var typeStr string
		if s, ok := t.(string); ok {
			typeStr = s
		} else {
			return fmt.Errorf("type entry doesn't point to a string entry, type: %s", name)
		}
		return extractNormalType(name, valuesMap, alreadyExtractedTypes, topLevel, typeStr)
	}
	return fmt.Errorf("missing type, ref or enum entry for type: %s", name)
}

func extractEnumType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any,
	topLevel bool, enumValues any) error {
	if i, ok := enumValues.([]int); ok {
		// int array
		newType := types.IntEnumType{
			Name:   name,
			Values: i,
		}
		alreadyExtractedTypes[name] = newType
	} else if s, ok := enumValues.([]string); ok {
		// string array
		newType := types.StringEnumType{
			Name:   name,
			Values: s,
		}
		alreadyExtractedTypes[name] = newType
	} else {
		return fmt.Errorf("unknown array entry for enum type with name: %s", name)
	}
	return nil
}

func extractRefType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any,
	topLevel bool, refStr string) error {
	return fmt.Errorf("TODO")
}

func extractNormalType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any,
	topLevel bool, typeStr string) error {
	return fmt.Errorf("TODO")
}
