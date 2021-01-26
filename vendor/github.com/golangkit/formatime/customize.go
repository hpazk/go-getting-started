package formatime

import (
	"fmt"
	"time"
)

// CustomTime is a nullable time.Time. It supports SQL, JSON serialization AND custom format.
// It will marshal to null if null.
type CustomTime struct {
	protocol
	Format string
}

// new custom format of time, s is timestamp
// examples:
// ANSIC       = "Mon Jan _2 15:04:05 2006"
// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// RFC822      = "02 Jan 06 15:04 MST"
// RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// RFC3339     = "2006-01-02T15:04:05Z07:00"
// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// Kitchen     = "3:04PM"
// // Handy time stamps.
// Stamp      = "Jan _2 15:04:05"
// StampMilli = "Jan _2 15:04:05.000"
// StampMicro = "Jan _2 15:04:05.000000"
// StampNano  = "Jan _2 15:04:05.000000000"
func NewCustomTime(s int64, format string) CustomTime {
	return CustomTime{
		protocol: protocol{
			Time:  time.Unix(s, 0),
			Valid: true,
		},
		Format: format,
	}
}

// new custom format of time from time.Time
func NewCustomTimeFrom(t time.Time, format string) CustomTime {
	return CustomTime{
		protocol: protocol{
			Time:  t,
			Valid: true,
		},
		Format: format,
	}
}

// new custom format of time right now
func NewCustomTimeNow(format string) CustomTime {
	return CustomTime{
		protocol: protocol{
			Time:  time.Now(),
			Valid: true,
		},
		Format: format,
	}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t CustomTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(t.Format))
	return []byte(stamp), nil
}
