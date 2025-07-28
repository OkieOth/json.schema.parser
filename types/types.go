package types

import (
	o "github.com/okieoth/goptional/v3"
)

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
	Name o.Optional[string]
}

type ObjectType struct {
	Name o.Optional[string]
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
}

// This type is only used as place holder while parsing references
type DummyType struct {
	Source string
	Name   string
}

type Property struct {
	Name         string
	ValueType    any
	Reference    o.Optional[string]
	ForeignKeyTo o.Optional[string]
	Description  string
}
