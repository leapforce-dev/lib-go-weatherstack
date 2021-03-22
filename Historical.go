package weatherstack

import (
	"fmt"
	"net/url"
	"time"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type HistoricalResponse struct {
	Request    Request            `json:"request"`
	Location   Location           `json:"location"`
	Current    CurrentWeather     `json:"current"`
	Historical map[string]Weather `json:"historical"`
}

type GetHistoricalWeatherConfig struct {
	Query     string
	StartDate civil.Date
	EndDate   *civil.Date
	Hourly    *Hourly
	Interval  *Interval
	Units     *Units
	Language  *string
}

func (service *Service) GetHistoricalWeather(config GetHistoricalWeatherConfig) (*HistoricalResponse, *errortools.Error) {
	values := url.Values{}

	startDate := utilities.DateToTime(config.StartDate)

	if config.EndDate == nil {
		values.Add("historical_date", startDate.Format(dateFormat))
	} else {
		endDate := utilities.DateToTime(*config.EndDate)

		if startDate.After(endDate) {
			return nil, errortools.ErrorMessage("StartDate must be smaller or equal to EndDate.")
		}

		maxEndDate := startDate.Add(time.Duration(MaxDaysPerCall-1) * 24 * time.Hour)

		if endDate.After(maxEndDate) {
			return nil, errortools.ErrorMessage("Maximum time frame of 60 days exceeded.")
		}

		values.Add("historical_date_start", startDate.Format(dateFormat))
		values.Add("historical_date_end", endDate.Format(dateFormat))
	}

	values.Add("query", config.Query)

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

	historicalResponse := HistoricalResponse{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("historical?%s", values.Encode())),
		ResponseModel: &historicalResponse,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &historicalResponse, nil
}
