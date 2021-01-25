package formatime

import (
	"fmt"
	"time"
)

// Timestamp is a nullable time.Time. It supports SQL, JSON serialization AND timestamp format.
// It will marshal to null if null.
type Timestamp struct {
	protocol
}

// new timestamp format of time
func NewTimestamp(s int64) Timestamp {
	return Timestamp{protocol{
		Time:  time.Unix(s, 0),
		Valid: true,
	}}
}

// new timestamp format of time right now
func NewTimestampNow() Timestamp {
	return Timestamp{protocol{
		Time:  time.Now(),
		Valid: true,
	}}
}

// new timestamp format of time from time.Time
func NewTimestampFrom(t time.Time) Timestamp {
	return Timestamp{protocol{
		Time:  t,
		Valid: true,
	}}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%d\"", t.Time.Unix())
	return []byte(stamp), nil
}
