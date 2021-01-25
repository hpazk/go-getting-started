package formatime

import (
	"fmt"
	"time"
)

// Minute is a nullable time.Time. It supports SQL, JSON serialization AND datetime format(2006-01-02 15).
// It will marshal to null if null.
type Minute struct {
	protocol
}

const _MinuteFormat = "2006-01-02 15:04"

// new Minute format of time, s is timestamp
func NewMinute(s int64) Minute {
	return Minute{protocol{
		Time:  time.Unix(s, 0),
		Valid: true,
	}}
}

// new Minute format of time right now
func NewMinuteNow() Minute {
	return Minute{protocol{
		Time:  time.Now(),
		Valid: true,
	}}
}

// new Minute format of time from time.Time
func NewMinuteFrom(t time.Time) Minute {
	return Minute{protocol{
		Time:  t,
		Valid: true,
	}}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Minute) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(_MinuteFormat))
	return []byte(stamp), nil
}
