package types

import (
	o "github.com/okieoth/goptional/v3"
)

type StringType struct {
	Default   o.Optional[string]
	Format    o.Optional[string]
	MinLength o.Optional[string]
	MaxLength o.Optional[string]
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
	Default o.Optional[string]
}

type DateType struct {
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DateTimeType struct {
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type TimeType struct {
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DurationType struct {
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
	Default          o.Optional[float32]
	Minimum          o.Optional[float32]
	ExclusiveMinimum o.Optional[float32]
	Maximum          o.Optional[float32]
	ExclusiveMaximum o.Optional[float32]
}

type BoolType struct {
	Name    o.Optional[string]
	Default o.Optional[bool]
}

type ByteType struct {
}

type ObjectType struct {
	Name o.Optional[string]
}

type ComplexType struct {
	Source      string
	TopLevel    bool
	Name        string
	Description string
	Properties  []Property
}

type ArrayType struct {
	Source      string
	Name        string
	Description string
	ValueType   any
}

type MapType struct {
	Source      string
	Name        string
	Description string
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
