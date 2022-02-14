package null

import "encoding/json"

// String is a nullable string.
type String struct {
	Value string
	Valid bool
}

// StringFrom creates a new, valid String from an existing value.
func StringFrom(s string) String {
	return String{
		Value: s,
		Valid: true,
	}
}

// NewNullString creates a new, null String.
func NewNullString() String {
	return String{
		Valid: false,
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte(`null`), nil
	}

	return json.Marshal(s.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(r []byte) error {
	var str string
	if err := json.Unmarshal(r, &str); err != nil {
		s.Valid = false
		return err
	}

	s.Value = str
	s.Valid = true
	return nil
}
