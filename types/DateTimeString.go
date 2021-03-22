package weatherstack

import (
	"strconv"
	"strings"
	"time"
)

const (
	dateTimeStringFormat string = "2006-01-02 15:04"
)

type DateTimeString time.Time

func (d *DateTimeString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(dateTimeStringFormat, unquoted)
	if err != nil {
		return err
	}

	*d = DateTimeString(_t)
	return nil
}

func (d *DateTimeString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeString) Value() time.Time {
	return time.Time(d)
}
