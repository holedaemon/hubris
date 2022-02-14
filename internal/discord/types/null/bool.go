package null

import "encoding/json"

// Bool is a nullable boolean.
type Bool struct {
	Value bool
	Valid bool
}

// BoolFrom creates a new, valid Bool from an existing value.
func BoolFrom(val bool) Bool {
	return Bool{
		Value: val,
		Valid: true,
	}
}

// NewNullBool creates a new, invalid Bool.
func NewNullBool() Bool {
	return Bool{
		Valid: false,
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte(`null`), nil
	}

	return json.Marshal(b.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Bool) UnmarshalJSON(r []byte) error {
	var boo bool
	if err := json.Unmarshal(r, &boo); err != nil {
		b.Valid = false
		return err
	}

	b.Value = boo
	b.Valid = true
	return nil
}
