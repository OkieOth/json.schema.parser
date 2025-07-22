package jsonschemaparser

import (
	"encoding/json"
	"fmt"

	"github.com/okieoth/jsonschemaparser/types"
)

func ParseBytes(input []byte) ([]types.ComplexType, error) {
	var parsedSchema map[string]any

	if err := json.Unmarshal(input, &parsedSchema); err != nil {
		return []types.ComplexType{}, fmt.Errorf("error while unmarshalling schema: %v", err)
	}

	var complexTypes []types.ComplexType

	if definitions, ok := parsedSchema["definitions"]; ok {
		definitionsMap, ok := definitions.(map[string]any)
		if !ok {
			return []types.ComplexType{}, fmt.Errorf("error while converting to definitions map")
		}
		ct, err := parseTypesFromDefinition(definitionsMap)
		if err != nil {
			return []types.ComplexType{}, fmt.Errorf("error while parsing types in the definitions section: %v", err)
		}
		complexTypes = ct
	}

	if _, ok := parsedSchema["properties"]; ok {
		// to level object
		ct, err := parseTopLevelType(parsedSchema)
		if err != nil {
			return []types.ComplexType{}, fmt.Errorf("error while parsing main type: %v", err)
		}
		complexTypes = append(complexTypes, ct)
	}

	return complexTypes, nil
}

func parseTopLevelType(parsedSchema map[string]any) (types.ComplexType, error) {
	return types.ComplexType{}, fmt.Errorf("TODO, not implemented yet")
}

func parseTypesFromDefinition(definitionsMap map[string]any) ([]types.ComplexType, error) {
	return nil, fmt.Errorf("TODO, not implemented yet")
}
