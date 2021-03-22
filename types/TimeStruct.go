package weatherstack

import (
	"strconv"
	"strings"
	"time"
)

const (
	timeStructFormat string = "03:04 PM"
)

type TimeStruct struct {
	TimeString string
	TimeTime   *time.Time
}

func (d *TimeStruct) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	unquoted = strings.Trim(unquoted, " ")

	if unquoted == "" {
		d = nil
		return nil
	}

	(*d).TimeString = unquoted

	_t, err := time.Parse(timeStringFormat, unquoted)
	if err == nil {
		(*d).TimeTime = &_t
	}

	return nil
}

func (d *TimeStruct) ValueString() *string {
	if d == nil {
		return nil
	}

	_d := string(d.TimeString)
	return &_d
}

func (d *TimeStruct) ValueTime() *time.Time {
	if d == nil {
		return nil
	}

	if d.TimeTime == nil {
		return nil
	}

	_d := time.Time(*d.TimeTime)
	return &_d
}
