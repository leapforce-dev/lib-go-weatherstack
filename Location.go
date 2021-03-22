package weatherstack

import (
	go_types "github.com/leapforce-libraries/go_types"
	w_types "github.com/leapforce-libraries/go_weatherstack/types"
)

type Location struct {
	Name           string                 `json:"name"`
	Country        string                 `json:"country"`
	Region         string                 `json:"region"`
	Lat            go_types.Float64String `json:"lat"`
	Lon            go_types.Float64String `json:"lon"`
	TimezoneID     string                 `json:"timezone_id"`
	Localtime      w_types.DateTimeString `json:"localtime"`
	LocaltimeEpoch int64                  `json:"localtime_epoch"`
	UTCOffset      go_types.Float64String `json:"utc_offset"`
}
