package jsonschemaparser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"
	"unicode"

	o "github.com/okieoth/goptional/v3"
	"github.com/okieoth/jsonschemaparser/types"
)

const NO_NAME = ""

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

	// to level object
	err := parseTopLevelType(parsedSchema, extractedTypes)
	if err != nil {
		return extractedTypes, fmt.Errorf("error while parsing main type: %v", err)
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
	_, err := extractType(typeName, parsedSchema, alreadyExtractedTypes, true)
	return err
}

func parseTypesFromDefinition(definitionsMap map[string]any, alreadyExtractedTypes map[string]any) error {
	for typeName, v := range definitionsMap {
		valuesMap, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("entry in definitions map, isn't a map type")
		}
		_, err := extractType(typeName, valuesMap, alreadyExtractedTypes, true)
		if err != nil {
			return fmt.Errorf("error while extracting type: %v", err)
		}
	}
	return nil
}

func extractType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (any, error) {
	if v, ok := valuesMap["enum"]; ok {
		// found enum entry
		return extractEnumType(name, alreadyExtractedTypes, v)
	} else if r, ok := valuesMap["$ref"]; ok {
		// found ref entry
		var refStr string
		if s, ok := r.(string); ok {
			refStr = s
		} else {
			return types.DummyType{}, fmt.Errorf("$ref doesn't point to a string entry, type: %s", name)
		}
		return extractRefType(name, valuesMap, alreadyExtractedTypes, topLevel, refStr)
	} else if t, ok := valuesMap["type"]; ok {
		// found type entry
		var typeStr string
		if s, ok := t.(string); ok {
			typeStr = s
		} else {
			return types.DummyType{}, fmt.Errorf("type entry doesn't point to a string entry, type: %s", name)
		}
		return extractNormalType(name, valuesMap, alreadyExtractedTypes, topLevel, typeStr)
	}
	return types.DummyType{}, fmt.Errorf("missing type, ref or enum entry for type: %s", name)
}

func toStringArray(a []any) []string {
	ret := make([]string, 0)
	for _, v := range a {
		if s, ok := v.(string); ok {
			ret = append(ret, s)
		}
	}
	return ret
}

func toIntArray(a []any) []int {
	ret := make([]int, 0)
	for _, v := range a {
		if s, ok := v.(int); ok {
			ret = append(ret, s)
		} else if f, ok := v.(float64); ok {
			ret = append(ret, int(f))
		}
	}
	return ret
}

func extractEnumType(name string, alreadyExtractedTypes map[string]any, enumValues any) (any, error) {
	if a, ok := enumValues.([]any); ok {
		if len(a) > 0 {
			if _, isInt := a[0].(int); isInt {
				newType := types.IntEnumType{
					Name:   name,
					Values: toIntArray(a),
				}
				alreadyExtractedTypes[name] = newType
				return newType, nil
			} else if _, isStr := a[0].(string); isStr {
				newType := types.StringEnumType{
					Name:   name,
					Values: toStringArray(a),
				}
				alreadyExtractedTypes[name] = newType
				return newType, nil
			} else if _, isFloat := a[0].(float64); isFloat {
				// int values are read as numbers by go ... that means float64
				newType := types.IntEnumType{
					Name:   name,
					Values: toIntArray(a),
				}
				alreadyExtractedTypes[name] = newType
				return newType, nil
			} else {
				return types.StringEnumType{}, fmt.Errorf("unknown array entry for enum type with name: %s, type: %v", name, reflect.TypeOf(a[0]))
			}
		} else {
			return types.StringEnumType{}, fmt.Errorf("enum array entry has len 0 for enum type with name: %s", name)
		}

	} else {
		return types.StringEnumType{}, fmt.Errorf("no array entry for enum type with name: %s", name)
	}
}

func extractRefType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any,
	topLevel bool, refStr string) (any, error) {
	return types.DummyType{}, fmt.Errorf("TODO")
}

func extractNormalType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any,
	topLevel bool, typeStr string) (any, error) {
	switch typeStr {
	case "integer":
		return extractIntegerType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "number":
		return extractNumberType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "boolean":
		return extractBooleanType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "string":
		return extractStringType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "array":
		return extractArrayType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "object":
		return extractObjectType(name, valuesMap, alreadyExtractedTypes, topLevel)
	default:
		return types.DummyType{}, fmt.Errorf("unknown type for name: %s, type: %s", name, typeStr)
	}
}

func getOptionalString(key string, valuesMap map[string]any, allowed []string) o.Optional[string] {
	if f, ok := valuesMap[key]; ok {
		if s, isStr := f.(string); isStr {
			if allowed == nil || slices.Contains(allowed, s) {
				return o.NewOptionalValue(s)
			}
		}
	}
	return o.NewOptional[string]()
}

func getOptionalInt(key string, valuesMap map[string]any, allowed []int) o.Optional[int] {
	if f, ok := valuesMap[key]; ok {
		if v, isStr := f.(float64); isStr { // needs to be float64, because JSON only now numbers by default
			s := int(v)
			if allowed == nil || slices.Contains(allowed, s) {
				return o.NewOptionalValue(s)
			}
		}
	}
	return o.NewOptional[int]()
}

func getOptionalNumber(key string, valuesMap map[string]any, allowed []float64) o.Optional[float64] {
	if f, ok := valuesMap[key]; ok {
		if v, isStr := f.(float64); isStr { // needs to be float64, because JSON only now numbers by default
			if allowed == nil || slices.Contains(allowed, v) {
				return o.NewOptionalValue(v)
			}
		}
	}
	return o.NewOptional[float64]()
}

func getOptionalBool(key string, valuesMap map[string]any) o.Optional[bool] {
	if f, ok := valuesMap[key]; ok {
		if v, isBool := f.(bool); isBool { // needs to be float64, because JSON only now numbers by default
			return o.NewOptionalValue(v)
		}
	}
	return o.NewOptional[bool]()
}

func extractIntegerType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.IntegerType, error) {
	intType := types.IntegerType{
		Name:             o.NewOptionalValue(name),
		Format:           getOptionalString("format", valuesMap, []string{"int32", "int64", "uint32", "uint64"}),
		Default:          getOptionalInt("default", valuesMap, nil),
		MultipleOf:       getOptionalInt("multipleOf", valuesMap, nil),
		Minimum:          getOptionalInt("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalInt("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalInt("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalInt("exclusiveMaximum", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return intType, fmt.Errorf("can't add int type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = intType
	}
	return intType, nil
}

func extractNumberType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.NumberType, error) {
	numberType := types.NumberType{
		Name:             o.NewOptionalValue(name),
		Format:           getOptionalString("format", valuesMap, []string{"float32", "float64"}),
		Default:          getOptionalNumber("default", valuesMap, nil),
		Minimum:          getOptionalNumber("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalNumber("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalNumber("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalNumber("exclusiveMaximum", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return numberType, fmt.Errorf("can't add float type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = numberType
	}
	return numberType, nil
}
func extractBooleanType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.BoolType, error) {
	boolType := types.BoolType{
		Name:    o.NewOptionalValue(name),
		Default: getOptionalBool("default", valuesMap),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return boolType, fmt.Errorf("can't add bool type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = boolType
	}
	return boolType, nil
}
func extractStringType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (any, error) {
	f := getOptionalString("format", valuesMap, nil)
	if formatValue, isSet := f.Get(); isSet {
		switch formatValue {
		case "date":
			return extractDateType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "time":
			return extractTimeType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "date-time":
			return extractDateTimeType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "uuid":
			return extractUuidType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "duration":
			return extractDurationType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "binary":
			return extractBinaryType(name, valuesMap, alreadyExtractedTypes, topLevel)
		}
	}
	return extractPureStringType(name, valuesMap, alreadyExtractedTypes, topLevel, f)
}

func extractDateType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.DateType, error) {
	t := types.DateType{
		Name:             o.NewOptionalValue(name),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add date type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractTimeType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.TimeType, error) {
	t := types.TimeType{
		Name:             o.NewOptionalValue(name),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add time type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractDateTimeType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.DateTimeType, error) {
	t := types.DateTimeType{
		Name:             o.NewOptionalValue(name),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add date-time type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractUuidType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.UUIDType, error) {
	t := types.UUIDType{
		Name:    o.NewOptionalValue(name),
		Default: getOptionalString("default", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add uuid type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractDurationType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.DurationType, error) {
	t := types.DurationType{
		Name:    o.NewOptionalValue(name),
		Default: getOptionalString("default", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add duration type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractBinaryType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.BinaryType, error) {
	t := types.BinaryType{
		Name: o.NewOptionalValue(name),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add binary type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractPureStringType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool, formatValue o.Optional[string]) (types.StringType, error) {
	t := types.StringType{
		Name:      o.NewOptionalValue(name),
		Default:   getOptionalString("default", valuesMap, nil),
		Format:    formatValue,
		MinLength: getOptionalInt("minLength", valuesMap, nil),
		MaxLength: getOptionalInt("minLength", valuesMap, nil),
		Pattern:   getOptionalString("pattern", valuesMap, nil),
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add string type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func getValueType(key string, valuesMap map[string]any, alreadyExtractedTypes map[string]any) (any, error) {
	if f, ok := valuesMap[key]; ok {
		if v, isMap := f.(map[string]any); isMap { // needs to be float64, because JSON only now numbers by default
			return extractType(NO_NAME, v, alreadyExtractedTypes, false)
		}
	} else {
		return types.DummyType{}, fmt.Errorf("couldn't find key to extract the value type")
	}
	return types.DummyType{}, nil // TODO
}

func extractArrayType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.ArrayType, error) {
	valueType, err := getValueType("items", valuesMap, alreadyExtractedTypes)
	if err != nil {
		return types.ArrayType{}, fmt.Errorf("error while extract value type (name: %s): %v", name, err)
	}
	t := types.ArrayType{
		Name:        o.NewOptionalValue(name),
		MinItems:    getOptionalInt("minItems", valuesMap, nil),
		MaxItems:    getOptionalInt("maxItems", valuesMap, nil),
		Description: getOptionalString("description", valuesMap, nil),
		ValueType:   valueType,
	}
	if name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes[name]
		if exist {
			return t, fmt.Errorf("can't add Array type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes[name] = t
	}
	return t, nil
}

func extractObjectType(name string, valuesMap map[string]any, alreadyExtractedTypes map[string]any, topLevel bool) (types.MapType, error) {
	return types.MapType{}, nil // TODO
}
