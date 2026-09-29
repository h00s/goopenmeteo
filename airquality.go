package goopenmeteo

import "encoding/json"

type AirQualityVariable = string
type AirQualityVariables = []string

const (
	// Air quality indices
	EuropeanAQI AirQualityVariable = "european_aqi"
	USAQI       AirQualityVariable = "us_aqi"

	// Pollutants
	PM10            AirQualityVariable = "pm10"
	PM2_5           AirQualityVariable = "pm2_5"
	CarbonMonoxide  AirQualityVariable = "carbon_monoxide"
	NitrogenDioxide AirQualityVariable = "nitrogen_dioxide"
	SulphurDioxide  AirQualityVariable = "sulphur_dioxide"
	Ozone           AirQualityVariable = "ozone"
	Dust            AirQualityVariable = "dust"
	UVIndex         AirQualityVariable = "uv_index"

	// Pollen (Europe only, during the pollen season)
	AlderPollen   AirQualityVariable = "alder_pollen"
	BirchPollen   AirQualityVariable = "birch_pollen"
	GrassPollen   AirQualityVariable = "grass_pollen"
	MugwortPollen AirQualityVariable = "mugwort_pollen"
	OlivePollen   AirQualityVariable = "olive_pollen"
	RagweedPollen AirQualityVariable = "ragweed_pollen"
)

type AirQuality struct {
	Latitude             float64           `json:"latitude"`
	Longitude            float64           `json:"longitude"`
	GenerationTimeMS     float64           `json:"generationTimeMs"`
	UTCOffsetSeconds     int               `json:"utcOffsetSeconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezoneAbbreviation"`
	Elevation            float64           `json:"elevation"`
	CurrentUnits         map[string]string `json:"currentUnits,omitempty"`
	Current              Current           `json:"current,omitempty"`
	HourlyUnits          map[string]string `json:"hourlyUnits,omitempty"`
	Hourly               TimeseriesData    `json:"hourly,omitempty"`
}

type AirQualityDomain = string

const (
	AirQualityDomainAuto       AirQualityDomain = "auto"
	AirQualityDomainCAMSEurope AirQualityDomain = "cams_europe"
	AirQualityDomainCAMSGlobal AirQualityDomain = "cams_global"
)

type AirQualityOptions struct {
	Latitude      float64             `url:"latitude"`
	Longitude     float64             `url:"longitude"`
	Current       AirQualityVariables `url:"current,omitempty"`
	Hourly        AirQualityVariables `url:"hourly,omitempty"`
	Domains       AirQualityDomain    `url:"domains,omitempty"`
	Timeformat    TimeFormat          `url:"timeformat,omitempty"`
	Timezone      string              `url:"timezone,omitempty"`
	PastDays      int                 `url:"past_days,omitempty"`
	ForecastDays  int                 `url:"forecast_days,omitempty"`
	ForecastHours int                 `url:"forecast_hours,omitempty"`
	PastHours     int                 `url:"past_hours,omitempty"`
	StartDate     string              `url:"start_date,omitempty"`
	EndDate       string              `url:"end_date,omitempty"`
	StartHour     string              `url:"start_hour,omitempty"`
	EndHour       string              `url:"end_hour,omitempty"`
	CellSelection CellSelection       `url:"cell_selection,omitempty"`
}

func (a *AirQualityOptions) ToQuery() string {
	return urlValues(a).Encode()
}

// UnmarshalJSON decodes an Open-Meteo air quality response; times are placed
// in the response's zone, as for Forecast.
func (a *AirQuality) UnmarshalJSON(data []byte) error {
	var w responseWire
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	loc := w.location()
	*a = AirQuality{
		Latitude:             w.Latitude,
		Longitude:            w.Longitude,
		GenerationTimeMS:     w.GenerationTimeMS,
		UTCOffsetSeconds:     w.UTCOffsetSeconds,
		Timezone:             w.Timezone,
		TimezoneAbbreviation: w.TimezoneAbbreviation,
		Elevation:            w.Elevation,
		CurrentUnits:         w.CurrentUnits,
		Current:              w.Current.in(loc),
		HourlyUnits:          w.HourlyUnits,
		Hourly:               w.Hourly.in(loc),
	}
	return nil
}
