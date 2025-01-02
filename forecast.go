package goopenmeteo

type Forecast struct {
	Latitude             float64           `json:"latitude"`
	Longitude            float64           `json:"longitude"`
	GenerationTimeMS     float64           `json:"generationTimeMs"`
	UTCOffsetSeconds     int               `json:"utcOffsetSeconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezoneAbbreviation"`
	Elevation            float64           `json:"elevation"`
	CurrentUnits         map[string]string `json:"currentUnits,omitempty"`
	Current              Current           `json:"current,omitempty"`
	Minutely15Units      map[string]string `json:"minutely15Units,omitempty"`
	Minutely15           TimeseriesData    `json:"minutely15,omitempty"`
	HourlyUnits          map[string]string `json:"hourlyUnits,omitempty"`
	Hourly               TimeseriesData    `json:"hourly,omitempty"`
	DailyUnits           map[string]string `json:"dailyUnits,omitempty"`
	Daily                TimeseriesData    `json:"daily,omitempty"`
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
	Latitude          float64           `url:"latitude"`
	Longitude         float64           `url:"longitude"`
	Elevation         float64           `url:"elevation,omitempty"`
	Current           WeatherVariables  `url:"current,omitempty"`
	Minutely15        WeatherVariables  `url:"minutely15,omitempty"`
	Hourly            WeatherVariables  `url:"hourly,omitempty"`
	Daily             WeatherVariables  `url:"daily,omitempty"`
	TemperatureUnit   TemperatureUnit   `url:"temperatureUnit,omitempty"`
	WindSpeedUnit     WindSpeedUnit     `url:"windSpeedUnit,omitempty"`
	PrecipitationUnit PrecipitationUnit `url:"precipitationUnit,omitempty"`
	Timeformat        TimeFormat        `url:"timeformat,omitempty"`
	Timezone          string            `url:"timezone,omitempty"`
	PastDays          int               `url:"pastDays,omitempty"`
	ForecastDays      int               `url:"forecastDays,omitempty"`
	ForecastHours     int               `url:"forecastHours,omitempty"`
	PastHours         int               `url:"pastHours,omitempty"`
	PastMinutely15    int               `url:"pastMinutely15,omitempty"`
	StartDate         string            `url:"startDate,omitempty"`
	EndDate           string            `url:"endDate,omitempty"`
	StartHour         string            `url:"startHour,omitempty"`
	EndHour           string            `url:"endHour,omitempty"`
	StartMinutely15   string            `url:"startMinutely15,omitempty"`
	EndMinutely15     string            `url:"endMinutely15,omitempty"`
	Models            []string          `url:"models,omitempty"`
	CellSelection     CellSelection     `url:"cellSelection,omitempty"`
}

func (f *ForecastOptions) ToQuery() string {
	return urlValues(f).Encode()
}
