package goopenmeteo

type Forecast struct {
	Latitude             float64           `json:"latitude"`
	Longitude            float64           `json:"longitude"`
	GenerationTimeMS     float64           `json:"generationtime_ms"`
	UTCOffsetSeconds     int               `json:"utc_offset_seconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezone_abbreviation"`
	Elevation            float64           `json:"elevation"`
	HourlyUnits          map[string]string `json:"hourly_units"`
	Hourly               Hourly            `json:"hourly"`
}

type ForecastOptions struct {
	Latitude     float64  `url:"latitude"`
	Longitude    float64  `url:"longitude"`
	ForecastDays int      `url:"forecast_days,omitempty"`
	Hourly       []string `url:"hourly,omitempty"`
}

func (f *ForecastOptions) Query() string {
	return urlValues(f).Encode()
}
