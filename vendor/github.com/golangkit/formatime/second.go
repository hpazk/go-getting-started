package formatime


import (
	"fmt"
	"time"
)

// Second is a nullable time.Time. It supports SQL, JSON serialization AND datetime format(2006-01-02 15:04:05).
// It will marshal to null if null.
type Second struct {
	protocol
}

const _SecondFormat = "2006-01-02 15:04:05"

// new Second format of time, s is timestamp
func NewSecond(s int64) Second {
	return Second{protocol{
		Time:  time.Unix(s, 0),
		Valid: true,
	}}
}

// new Second format of time right now
func NewSecondNow() Second {
	return Second{protocol{
		Time:  time.Now(),
		Valid: true,
	}}
}

// new Second format of time from time.Time
func NewSecondFrom(t time.Time) Second {
	return Second{protocol{
		Time:  t,
		Valid: true,
	}}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Second) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(_SecondFormat))
	return []byte(stamp), nil
}
