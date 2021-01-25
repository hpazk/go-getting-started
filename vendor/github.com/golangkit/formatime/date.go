package formatime


import (
	"fmt"
	"time"
)

// Date is a nullable time.Time. It supports SQL, JSON serialization AND datetime format(2006-01-02).
// It will marshal to null if null.
type Date struct {
	protocol
}

const _DateFormat = "2006-01-02"

// new Date format of time, s is timestamp
func NewDate(s int64) Date {
	return Date{protocol{
		Time:  time.Unix(s, 0),
		Valid: true,
	}}
}

// new Date format of time right now
func NewDateNow() Date {
	return Date{protocol{
		Time:  time.Now(),
		Valid: true,
	}}
}

// new Date format of time from time.Time
func NewDateFrom(t time.Time) Date {
	return Date{protocol{
		Time:  t,
		Valid: true,
	}}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Date) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(_DateFormat))
	return []byte(stamp), nil
}
