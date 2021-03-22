package weatherstack

type Hourly int64

const (
	HourlyOn  Hourly = 1
	HourlyOff Hourly = 0
)

type Interval int64

const (
	Interval1Hour      Interval = 1
	Interval3Hours     Interval = 3
	Interval6Hours     Interval = 6
	IntervalDayNight   Interval = 12
	IntervalDayAverage Interval = 24
)

type Units string

const (
	UnitsMetric     Units = "m"
	UnitsScientific Units = "s"
	UnitsFahrenheit Units = "f"
)
