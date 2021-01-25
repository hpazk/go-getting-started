package formatime

import (
	"fmt"
	"time"
)

// Hour is a nullable time.Time. It supports SQL, JSON serialization AND datetime format(2006-01-02 15).
// It will marshal to null if null.
type Hour struct {
	protocol
}

const _HourFormat = "2006-01-02 15"

// new Hour format of time, s is timestamp
func NewHour(s int64) Hour {
	return Hour{protocol{
		Time:  time.Unix(s, 0),
		Valid: true,
	}}
}

// new Hour format of time right now
func NewHourNow() Hour {
	return Hour{protocol{
		Time:  time.Now(),
		Valid: true,
	}}
}

// new Hour format of time from time.Time
func NewHourFrom(t time.Time) Hour {
	return Hour{protocol{
		Time:  t,
		Valid: true,
	}}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Hour) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(_HourFormat))
	return []byte(stamp), nil
}
