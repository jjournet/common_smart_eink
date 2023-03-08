package common_smart_eink

import (
	"github.com/briandowns/openweathermap"
)

type Current_weather_storage struct {
	Zip      string                            `json:"Zip"`
	Dt       int                               `json:"Dt"`
	Expires  int64                             `json:"Expires"`
	OpenData openweathermap.OneCallCurrentData `json:"OpenData"`
}

type Hourly_weather_storage struct {
	Zip      string                           `json:"Zip"`
	Dt       int                              `json:"Dt"`
	Expires  int64                            `json:"Expires"`
	OpenData openweathermap.OneCallHourlyData `json:"OpenData"`
}

type Daily_weather_storage struct {
	Zip      string                          `json:"Zip"`
	Dt       int                             `json:"Dt"`
	Expires  int64                           `json:"Expires"`
	OpenData openweathermap.OneCallDailyData `json:"OpenData"`
}

type Alert_weather_storage struct {
	Zip      string                            `json:"Zip"`
	Dt       int                               `json:"Dt"`
	Expires  int64                             `json:"Expires"`
	Latest   int                               `json:"Latest"`
	OpenData []openweathermap.OneCallAlertData `json:"OpenData"`
}
