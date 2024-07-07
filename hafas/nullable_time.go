package hafas

import "time"

type NullableTime struct {
	time.Time
}

func (n NullableTime) MarshalJSON() ([]byte, error) {
	if n.IsZero() {
		return []byte("null"), nil
	} else {
		return n.Time.MarshalJSON()
	}
}
