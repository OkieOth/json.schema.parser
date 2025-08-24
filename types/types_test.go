package types_test

import (
	"testing"

	"github.com/okieoth/jsonschemaparser/types"
	"github.com/stretchr/testify/require"
)

func TestNewParseResult(t *testing.T) {
	r := types.NewParseResult()
	require.NotNil(t, r.ArrayTypes)
	require.NotNil(t, r.MapTypes)
	require.NotNil(t, r.ComplexTypes)
	require.NotNil(t, r.IntEnums)
	require.NotNil(t, r.StringEnums)
	require.NotNil(t, r.IntegerTypes)
	require.NotNil(t, r.NumberTypes)
	require.NotNil(t, r.StringTypes)
	require.NotNil(t, r.UUIDTypes)
	require.NotNil(t, r.DateTypes)
	require.NotNil(t, r.DateTimeTypes)
	require.NotNil(t, r.TimeTypes)
	require.NotNil(t, r.DurationTypes)
	require.NotNil(t, r.BoolTypes)
	require.NotNil(t, r.BinaryTypes)
	require.NotNil(t, r.ObjectTypes)

	require.Equal(t, 0, len(r.ArrayTypes))
	require.Equal(t, 0, len(r.MapTypes))
	require.Equal(t, 0, len(r.ComplexTypes))
	require.Equal(t, 0, len(r.IntEnums))
	require.Equal(t, 0, len(r.StringEnums))
	require.Equal(t, 0, len(r.IntegerTypes))
	require.Equal(t, 0, len(r.NumberTypes))
	require.Equal(t, 0, len(r.StringTypes))
	require.Equal(t, 0, len(r.UUIDTypes))
	require.Equal(t, 0, len(r.DateTypes))
	require.Equal(t, 0, len(r.DateTimeTypes))
	require.Equal(t, 0, len(r.TimeTypes))
	require.Equal(t, 0, len(r.DurationTypes))
	require.Equal(t, 0, len(r.BoolTypes))
	require.Equal(t, 0, len(r.BinaryTypes))
	require.Equal(t, 0, len(r.ObjectTypes))
}
