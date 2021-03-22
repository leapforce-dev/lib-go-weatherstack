package weatherstack

import (
	w_types "github.com/leapforce-libraries/go_weatherstack/types"
)

type CurrentWeather struct {
	ObservationTime     w_types.TimeString `json:"observation_time"`
	Temperature         int64              `json:"temperature"`
	WeatherCode         int64              `json:"weather_code"`
	WeatherIcons        []string           `json:"weather_icons"`
	WeatherDescriptions []string           `json:"weather_descriptions"`
	WindSpeed           int64              `json:"wind_speed"`
	WindDegree          int64              `json:"wind_degree"`
	WindDir             string             `json:"wind_dir"`
	Pressure            int64              `json:"pressure"`
	Precip              float64            `json:"precip"`
	Humidity            int64              `json:"humidity"`
	Cloudcover          int64              `json:"cloudcover"`
	FeelsLike           int64              `json:"feelslike"`
	UVIndex             int64              `json:"uv_index"`
	Visibility          int64              `json:"visibility"`
	IsDay               string             `json:"is_day"`
}
