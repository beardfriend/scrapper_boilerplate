package types

import (
	"database/sql/driver"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type DateString time.Time

const DateFormat = "2006-01-02"

func (t DateString) Parse(layout string, value string) (DateString, error) {
	time, err := time.Parse(layout, value)
	if err != nil {
		return DateString{}, err
	}

	return DateString(time), nil
}

func (t *DateString) ToString() string {
	return time.Time(*t).Format(DateFormat)
}

func (t *DateString) ToTime() time.Time {
	return time.Time(*t)
}

func (t DateString) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t *DateString) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	parsed, _ := time.Parse(DateFormat, str)
	*t = DateString(parsed)
	return nil
}

func (t DateString) Now() DateString {
	return DateString(t)
}

func (t DateString) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *DateString) Scan(value interface{}) error {
	if value == nil {
		*t = DateString(time.Now())
		return nil
	}
	if v, ok := value.(time.Time); ok {
		*t = DateString(v)
		return nil
	}

	return errors.New("cannot scan")
}
