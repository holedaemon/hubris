package null

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

const (
	strTest = "testing null strings lol"
)

type typeTest struct {
	String String
	Int    Int
	Int64  Int64
	Bool   Bool
}

func TestTypes(t *testing.T) {
	ts := StringFrom(strTest)
	assert.Assert(t, ts.Valid && ts.Value == strTest, "string has unexpected values")

	tb := BoolFrom(false)
	assert.Assert(t, tb.Valid && !tb.Value, "bool has unexpected values")

	ti := IntFrom(69)
	assert.Assert(t, ti.Valid && ti.Value == 69, "int has unexpected values")

	tib := Int64From(42069)
	assert.Assert(t, tib.Valid && tib.Value == 42069, "int64 has unexpected values")

	tt := typeTest{
		String: ts,
		Int:    ti,
		Int64:  tib,
		Bool:   tb,
	}

	js, err := json.MarshalIndent(tt, "", "	")
	assert.NilError(t, err, "marshaling json")

	t.Log("\n" + string(js))

	ts = NewNullString()
	assert.Assert(t, !ts.Valid, "null string is valid")

	tb = NewNullBool()
	assert.Assert(t, !tb.Valid, "null bool is valid")

	ti = NewNullInt()
	assert.Assert(t, !ti.Valid, "null int is valid")

	tib = NewNullInt64()
	assert.Assert(t, !tib.Valid, "null int64 is valid")

	tt = typeTest{
		String: ts,
		Int:    ti,
		Int64:  tib,
		Bool:   tb,
	}

	js, err = json.MarshalIndent(tt, "", "	")
	assert.NilError(t, err, "marshaling json")

	t.Log("\n" + string(js))
}
