package formatime

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// protocol is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
type protocol struct {
	Time  time.Time
	Valid bool
}

// Scan implements the Scanner interface.
func (t *protocol) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = x
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	t.Valid = err == nil
	return err
}

// Value implements the driver Valuer interface.
func (t protocol) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// IsZero returns true for invalid Times, hopefully for future omitempty support.
// A non-null Time with a zero value will not be considered zero.
func (t protocol) IsZero() bool {
	return !t.Valid
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (t *protocol) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		err = t.Time.UnmarshalJSON(data)
	case map[string]interface{}:
		ti, tiOK := x["Time"].(string)
		valid, validOK := x["Valid"].(bool)
		if !tiOK || !validOK {
			return fmt.Errorf(`json: unmarshalling object into Go value of type null.Time requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
		}
		err = t.Time.UnmarshalText([]byte(ti))
		t.Valid = valid
		return err
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

// Marshal
func (t protocol) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalText()
}

func (t *protocol) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	if err := t.Time.UnmarshalText(text); err != nil {
		return err
	}
	t.Valid = true
	return nil
}
