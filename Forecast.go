package weatherstack

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ForecastResponse struct {
	Request  Request            `json:"request"`
	Location Location           `json:"location"`
	Current  CurrentWeather     `json:"current"`
	Forecast map[string]Weather `json:"forecast"`
}

type GetForecastWeatherConfig struct {
	Query        string
	ForecastDays *uint
	Hourly       *Hourly
	Interval     *Interval
	Units        *Units
	Language     *string
}

func (service *Service) GetForecastWeather(config GetForecastWeatherConfig) (*ForecastResponse, *errortools.Error) {
	values := url.Values{}

	values.Add("query", config.Query)

	if config.ForecastDays != nil {
		values.Add("forecast_days", fmt.Sprintf("%v", *config.ForecastDays))
	}

	if config.Hourly != nil {
		values.Add("hourly", fmt.Sprintf("%v", int64(*config.Hourly)))
	}

	if config.Interval != nil {
		values.Add("interval", fmt.Sprintf("%v", int64(*config.Interval)))
	}

	if config.Units != nil {
		values.Add("units", fmt.Sprintf("%s", string(*config.Units)))
	}

	if config.Language != nil {
		values.Add("language", *config.Language)
	}

	forecastResponse := ForecastResponse{}

	requestConfig := go_http.RequestConfig{
		Url:           service.url(fmt.Sprintf("forecast?%s", values.Encode())),
		ResponseModel: &forecastResponse,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &forecastResponse, nil
}
