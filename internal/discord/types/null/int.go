package null

import "encoding/json"

// Int is a nullable int.
type Int struct {
	Value int
	Valid bool
}

// IntFrom creates a new, valid Int from an existing value.
func IntFrom(i int) Int {
	return Int{
		Value: i,
		Valid: true,
	}
}

// NewNullInt creates a new, null Int.
func NewNullInt() Int {
	return Int{
		Valid: false,
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte(`null`), nil
	}

	return json.Marshal(i.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int) UnmarshalJSON(r []byte) error {
	var in int
	if err := json.Unmarshal(r, &in); err != nil {
		i.Valid = false
		return err
	}

	i.Value = in
	i.Valid = true
	return nil
}
