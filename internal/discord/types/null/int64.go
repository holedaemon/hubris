package null

import "encoding/json"

// Int64 is a nullable int64.
type Int64 struct {
	Value int64
	Valid bool
}

// Int64From creates a new, valid Int from an existing value.
func Int64From(i int64) Int64 {
	return Int64{
		Value: i,
		Valid: true,
	}
}

// NewNullInt64 creates a new, null Int.
func NewNullInt64() Int64 {
	return Int64{
		Valid: false,
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte(`null`), nil
	}

	return json.Marshal(i.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(r []byte) error {
	var in int64
	if err := json.Unmarshal(r, &in); err != nil {
		i.Valid = false
		return err
	}

	i.Value = in
	i.Valid = true
	return nil
}
