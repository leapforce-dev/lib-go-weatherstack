package weatherstack

import (
	"strconv"
	"strings"
	"time"
)

const (
	timeStringFormat string = "03:04 PM"
)

type TimeString time.Time

func (d *TimeString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	unquoted = strings.Trim(unquoted, " ")

	if unquoted == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(timeStringFormat, unquoted)
	if err != nil {
		return err
	}

	*d = TimeString(_t)
	return nil
}

func (d *TimeString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d TimeString) Value() time.Time {
	return time.Time(d)
}
