package types

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(r []byte) error {
	var ts time.Time
	if err := ts.UnmarshalJSON(r); err != nil {
		return err
	}

	t.Time = ts
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, t.Format(time.RFC3339))), nil
}

func Now() Time {
	return Time{
		Time: time.Now(),
	}
}

func TimeToTime(t time.Time) Time {
	return Time{
		Time: t,
	}
}
