package types

import (
	o "github.com/okieoth/goptional/v3"
)

// This type only covers the main parsed types, that most likely need to be handled
// in the first place.
type ParseResult struct {
	ComplexTypes  map[string]ComplexType
	ArrayTypes    map[string]ArrayType
	MapTypes      map[string]MapType
	IntEnums      map[string]IntEnumType
	StringEnums   map[string]StringEnumType
	IntegerTypes  map[string]IntegerType
	NumberTypes   map[string]NumberType
	StringTypes   map[string]StringType
	UUIDTypes     map[string]UUIDType
	DateTypes     map[string]DateType
	DateTimeTypes map[string]DateTimeType
	TimeTypes     map[string]TimeType
	DurationTypes map[string]DurationType
	BoolTypes     map[string]BoolType
	BinaryTypes   map[string]BinaryType
	ObjectTypes   map[string]ObjectType
}

func NewParseResult() ParseResult {
	return ParseResult{
		ComplexTypes:  make(map[string]ComplexType, 0),
		ArrayTypes:    make(map[string]ArrayType, 0),
		MapTypes:      make(map[string]MapType, 0),
		IntEnums:      make(map[string]IntEnumType, 0),
		StringEnums:   make(map[string]StringEnumType, 0),
		IntegerTypes:  make(map[string]IntegerType, 0),
		NumberTypes:   make(map[string]NumberType, 0),
		StringTypes:   make(map[string]StringType, 0),
		UUIDTypes:     make(map[string]UUIDType, 0),
		DateTypes:     make(map[string]DateType, 0),
		DateTimeTypes: make(map[string]DateTimeType, 0),
		TimeTypes:     make(map[string]TimeType, 0),
		DurationTypes: make(map[string]DurationType, 0),
		BoolTypes:     make(map[string]BoolType, 0),
		BinaryTypes:   make(map[string]BinaryType, 0),
		ObjectTypes:   make(map[string]ObjectType, 0),
	}
}

type StringType struct {
	Name      o.Optional[string]
	Default   o.Optional[string]
	Format    o.Optional[string]
	MinLength o.Optional[int]
	MaxLength o.Optional[int]
	Pattern   o.Optional[string]
}

type IntEnumType struct {
	Name    string
	Default o.Optional[int]
	Values  []int
}

type StringEnumType struct {
	Name    string
	Default o.Optional[string]
	Values  []string
}

type UUIDType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

type DateType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DateTimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type TimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DurationType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

type IntegerType struct {
	Name             o.Optional[string]
	Format           o.Optional[string]
	Default          o.Optional[int]
	MultipleOf       o.Optional[int]
	Minimum          o.Optional[int]
	ExclusiveMinimum o.Optional[int]
	Maximum          o.Optional[int]
	ExclusiveMaximum o.Optional[int]
}

type NumberType struct {
	Name             o.Optional[string]
	Format           o.Optional[string]
	Default          o.Optional[float64]
	Minimum          o.Optional[float64]
	ExclusiveMinimum o.Optional[float64]
	Maximum          o.Optional[float64]
	ExclusiveMaximum o.Optional[float64]
}

type BoolType struct {
	Name    o.Optional[string]
	Default o.Optional[bool]
}

type BinaryType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

type ObjectType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

type ComplexType struct {
	Source      string
	TopLevel    bool
	Name        string
	Description o.Optional[string]
	Properties  []Property
}

type ArrayType struct {
	Source      string
	Name        o.Optional[string]
	MinItems    o.Optional[int]
	MaxItems    o.Optional[int]
	Description o.Optional[string]
	ValueType   any
}

type MapType struct {
	Source      string
	Name        string
	Description o.Optional[string]
	ValueType   any
	TopLevel    bool
}

// This type is only used as place holder while parsing references
type DummyType struct {
	Source string
	Name   string
}

type Property struct {
	Name         string
	ValueType    any
	ForeignKeyTo o.Optional[string]
	Description  o.Optional[string]
}
