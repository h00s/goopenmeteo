package goopenmeteo

type Forecast struct {
	Latitude             float64           `json:"latitude"`
	Longitude            float64           `json:"longitude"`
	GenerationTimeMS     float64           `json:"generationtime_ms"`
	UTCOffsetSeconds     int               `json:"utc_offset_seconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezone_abbreviation"`
	Elevation            float64           `json:"elevation"`
	CurrentUnits         map[string]string `json:"current_units"`
	Current              Current           `json:"current"`
	HourlyUnits          map[string]string `json:"hourly_units"`
	Hourly               Hourly            `json:"hourly"`
}

type TimeFormat = string
type TemperatureUnit = string
type WindSpeedUnit = string
type PrecipitationUnit = string
type CellSelection = string

const (
	TimeFormatISO8601  TimeFormat = "iso8601"
	TimeFormatUnixtime TimeFormat = "unixtime"

	TemperatureUnitCelsius    TemperatureUnit = "celsius"
	TemperatureUnitFahrenheit TemperatureUnit = "fahrenheit"

	WindSpeedUnitKMH WindSpeedUnit = "kmh"
	WindSpeedUnitMS  WindSpeedUnit = "ms"
	WindSpeedUnitMPH WindSpeedUnit = "mph"
	WindSpeedUnitKN  WindSpeedUnit = "kn"

	PrecipitationUnitMM   PrecipitationUnit = "mm"
	PrecipitationUnitInch PrecipitationUnit = "inch"

	CellSelectionLand    CellSelection = "land"
	CellSelectionSea     CellSelection = "sea"
	CellSelectionNearest CellSelection = "nearest"
)

type ForecastOptions struct {
	Latitude           float64           `url:"latitude"`
	Longitude          float64           `url:"longitude"`
	Elevation          float64           `url:"elevation,omitempty"`
	Hourly             WeatherVariables  `url:"hourly,omitempty"`
	Daily              WeatherVariables  `url:"daily,omitempty"`
	Current            WeatherVariables  `url:"current,omitempty"`
	TemperatureUnit    TemperatureUnit   `url:"temperature_unit,omitempty"`
	WindSpeedUnit      WindSpeedUnit     `url:"wind_speed_unit,omitempty"`
	PrecipitationUnit  PrecipitationUnit `url:"precipitation_unit,omitempty"`
	Timeformat         TimeFormat        `url:"timeformat,omitempty"`
	Timezone           string            `url:"timezone,omitempty"`
	PastDays           int               `url:"past_days,omitempty"`
	ForecastDays       int               `url:"forecast_days,omitempty"`
	ForecastHours      int               `url:"forecast_hours,omitempty"`
	PastHours          int               `url:"past_hours,omitempty"`
	ForecastMinutely15 int               `url:"forecast_minutely_15,omitempty"`
	PastMinutely15     int               `url:"past_minutely_15,omitempty"`
	StartDate          string            `url:"start_date,omitempty"`
	EndDate            string            `url:"end_date,omitempty"`
	StartHour          string            `url:"start_hour,omitempty"`
	EndHour            string            `url:"end_hour,omitempty"`
	StartMinutely15    string            `url:"start_minutely_15,omitempty"`
	EndMinutely15      string            `url:"end_minutely_15,omitempty"`
	Models             []string          `url:"models,omitempty"`
	CellSelection      CellSelection     `url:"cell_selection,omitempty"`
}

func (f *ForecastOptions) Query() string {
	return urlValues(f).Encode()
}
