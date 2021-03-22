package weatherstack

import (
	go_types "github.com/leapforce-libraries/go_types"
	w_types "github.com/leapforce-libraries/go_weatherstack/types"
)

type Weather struct {
	Date      w_types.DateString `json:"date"`
	DateEpoch int64              `json:"date_epoch"`
	Astro     Astro              `json:"astro"`
	MinTemp   int64              `json:"mintemp"`
	MaxTemp   int64              `json:"maxtemp"`
	AvgTemp   int64              `json:"avgtemp"`
	TotalSnow float64            `json:"totalsnow"`
	SunHour   float64            `json:"sunhour"`
	UVIndex   int64              `json:"uv_index"`
	Hourly    []HourlyWeather    `json:"hourly"`
}

type Astro struct {
	Sunrise          w_types.TimeString `json:"sunrise"`
	Sunset           w_types.TimeString `json:"sunset"`
	Moonrise         w_types.TimeString `json:"moonrise"`
	Moonset          w_types.TimeString `json:"moonset"`
	MoonPhase        string             `json:"moon_phase"`
	MoonIllumination int64              `json:"moon_illumination"`
}

type HourlyWeather struct {
	Time                go_types.Int64String `json:"time"`
	Temperature         int64                `json:"temperature"`
	WindSpeed           int64                `json:"wind_speed"`
	WindDegree          int64                `json:"wind_degree"`
	WindDir             string               `json:"wind_dir"`
	WeatherCode         int64                `json:"weather_code"`
	WeatherIcons        []string             `json:"weather_icons"`
	WeatherDescriptions []string             `json:"weather_descriptions"`
	Precip              float64              `json:"precip"`
	Humidity            int64                `json:"humidity"`
	Visibility          int64                `json:"visibility"`
	Pressure            int64                `json:"pressure"`
	Cloudcover          int64                `json:"cloudcover"`
	Heatindex           int64                `json:"heatindex"`
	Dewpoint            int64                `json:"dewpoint"`
	Windchill           int64                `json:"windchill"`
	Windgust            int64                `json:"windgust"`
	FeelsLike           int64                `json:"feelslike"`
	ChanceOfRain        int64                `json:"chanceofrain"`
	ChanceOfRemDry      int64                `json:"chanceofremdry"`
	ChanceOfWindy       int64                `json:"chanceofwindy"`
	ChanceOfOvercast    int64                `json:"chanceofovercast"`
	ChanceOfSunshine    int64                `json:"chanceofsunshine"`
	ChanceOfFrost       int64                `json:"chanceoffrost"`
	ChanceOfHighTemp    int64                `json:"chanceofhightemp"`
	ChanceOfFog         int64                `json:"chanceoffog"`
	ChanceOfSnow        int64                `json:"chanceofsnow"`
	ChanceOfThunder     int64                `json:"chanceofthunder"`
	UVIndex             int64                `json:"uv_index"`
}
