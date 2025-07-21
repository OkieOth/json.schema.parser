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
	Format           o.Optional[string]
	Default          o.Optional[int]
	Minimum          o.Optional[int]
	ExclusiveMinimum o.Optional[int]
	Maximum          o.Optional[int]
	ExclusiveMaximum o.Optional[int]
}

type NumberType struct {
	Format           o.Optional[string]
	Default          o.Optional[float32]
	Minimum          o.Optional[float32]
	ExclusiveMinimum o.Optional[float32]
	Maximum          o.Optional[float32]
	ExclusiveMaximum o.Optional[float32]
}

type BoolType struct {
	Default o.Optional[bool]
}

type ByteType struct {
}

type ObjectType struct {
}

type ComplexType struct {
	Source     string
	Name       string
	Properties []Property
}

type ArrayType struct {
	Source    string
	Name      o.Optional[string]
	ValueType any
}

type MapType struct {
	Source    string
	Name      o.Optional[string]
	ValueType any
}

type Property struct {
	Name         string
	ValueType    any
	Reference    o.Optional[string]
	ForeignKeyTo o.Optional[string]
}
