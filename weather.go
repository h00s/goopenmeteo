package goopenmeteo

type WeatherVariable = string
type WeatherVariables = []string

const (
	// Temperature and Comfort
	Temperature2M       WeatherVariable = "temperature_2m"
	RelativeHumidity2M  WeatherVariable = "relative_humidity_2m"
	DewPoint2M          WeatherVariable = "dew_point_2m"
	ApparentTemperature WeatherVariable = "apparent_temperature"
	// Daily Temperature
	Temperature2MMax       WeatherVariable = "temperature_2m_max"
	Temperature2MMin       WeatherVariable = "temperature_2m_min"
	ApparentTemperatureMax WeatherVariable = "apparent_temperature_max"
	ApparentTemperatureMin WeatherVariable = "apparent_temperature_min"

	// Pressure
	PressureMsl     WeatherVariable = "pressure_msl"
	SurfacePressure WeatherVariable = "surface_pressure"

	// Cloud Coverage
	CloudCover     WeatherVariable = "cloud_cover"
	CloudCoverLow  WeatherVariable = "cloud_cover_low"
	CloudCoverMid  WeatherVariable = "cloud_cover_mid"
	CloudCoverHigh WeatherVariable = "cloud_cover_high"

	// Wind
	WindSpeed10M      WeatherVariable = "wind_speed_10m"
	WindSpeed80M      WeatherVariable = "wind_speed_80m"
	WindSpeed120M     WeatherVariable = "wind_speed_120m"
	WindSpeed180M     WeatherVariable = "wind_speed_180m"
	WindDirection10M  WeatherVariable = "wind_direction_10m"
	WindDirection80M  WeatherVariable = "wind_direction_80m"
	WindDirection120M WeatherVariable = "wind_direction_120m"
	WindDirection180M WeatherVariable = "wind_direction_180m"
	WindGusts10M      WeatherVariable = "wind_gusts_10m"
	// Daily Wind
	WindSpeed10MMax          WeatherVariable = "wind_speed_10m_max"
	WindGusts10MMax          WeatherVariable = "wind_gusts_10m_max"
	WindDirection10MDominant WeatherVariable = "wind_direction_10m_dominant"

	// Solar Radiation
	ShortwaveRadiation     WeatherVariable = "shortwave_radiation"
	DirectRadiation        WeatherVariable = "direct_radiation"
	DirectNormalIrradiance WeatherVariable = "direct_normal_irradiance"
	DiffuseRadiation       WeatherVariable = "diffuse_radiation"
	GlobalTiltedIrradiance WeatherVariable = "global_tilted_irradiance"
	// Daily Solar
	ShortwaveRadiationSum WeatherVariable = "shortwave_radiation_sum"
	SunshineDuration      WeatherVariable = "sunshine_duration"
	DaylightDuration      WeatherVariable = "daylight_duration"

	// Moisture and Evaporation
	VapourPressureDeficit    WeatherVariable = "vapour_pressure_deficit"
	Cape                     WeatherVariable = "cape"
	Evapotranspiration       WeatherVariable = "evapotranspiration"
	ET0FaoEvapotranspiration WeatherVariable = "et0_fao_evapotranspiration"

	// Precipitation
	Precipitation            WeatherVariable = "precipitation"
	Snowfall                 WeatherVariable = "snowfall"
	PrecipitationProbability WeatherVariable = "precipitation_probability"
	Rain                     WeatherVariable = "rain"
	Showers                  WeatherVariable = "showers"
	// Daily Precipitation
	PrecipitationSum             WeatherVariable = "precipitation_sum"
	RainSum                      WeatherVariable = "rain_sum"
	ShowersSum                   WeatherVariable = "showers_sum"
	SnowfallSum                  WeatherVariable = "snowfall_sum"
	PrecipitationHours           WeatherVariable = "precipitation_hours"
	PrecipitationProbabilityMax  WeatherVariable = "precipitation_probability_max"
	PrecipitationProbabilityMin  WeatherVariable = "precipitation_probability_min"
	PrecipitationProbabilityMean WeatherVariable = "precipitation_probability_mean"

	// UV Index
	UVIndexMax         WeatherVariable = "uv_index_max"
	UVIndexClearSkyMax WeatherVariable = "uv_index_clear_sky_max"

	// Sun Times
	Sunrise WeatherVariable = "sunrise"
	Sunset  WeatherVariable = "sunset"

	// Other Atmospheric
	WeatherCode         WeatherVariable = "weather_code"
	SnowDepth           WeatherVariable = "snow_depth"
	FreezingLevelHeight WeatherVariable = "freezing_level_height"
	Visibility          WeatherVariable = "visibility"

	// Soil Conditions
	SoilTemperature0cm   WeatherVariable = "soil_temperature_0cm"
	SoilTemperature6cm   WeatherVariable = "soil_temperature_6cm"
	SoilTemperature18cm  WeatherVariable = "soil_temperature_18cm"
	SoilTemperature54cm  WeatherVariable = "soil_temperature_54cm"
	SoilMoisture0To1cm   WeatherVariable = "soil_moisture_0_to_1cm"
	SoilMoisture1To3cm   WeatherVariable = "soil_moisture_1_to_3cm"
	SoilMoisture3To9cm   WeatherVariable = "soil_moisture_3_to_9cm"
	SoilMoisture9To27cm  WeatherVariable = "soil_moisture_9_to_27cm"
	SoilMoisture27To81cm WeatherVariable = "soil_moisture_27_to_81cm"

	// Day/Night
	IsDay WeatherVariable = "is_day"
)

var WeatherCodes = map[int]string{
	0:  "Clear sky",
	1:  "Mainly clear",
	2:  "Partly cloudy",
	3:  "Overcast",
	45: "Fog",
	48: "Depositing rime fog",
	51: "Light drizzle",
	53: "Moderate drizzle",
	55: "Dense drizzle",
	56: "Light freezing drizzle",
	57: "Dense freezing drizzle",
	61: "Slight rain",
	63: "Moderate rain",
	65: "Heavy rain",
	66: "Light freezing rain",
	67: "Heavy freezing rain",
	71: "Slight snow fall",
	73: "Moderate snow fall",
	75: "Heavy snow fall",
	77: "Snow grains",
	80: "Slight rain showers",
	81: "Moderate rain showers",
	82: "Violent rain showers",
	85: "Slight snow showers",
	86: "Heavy snow showers",
	95: "Thunderstorm",
	96: "Thunderstorm with slight hail",
	99: "Thunderstorm with heavy hail",
}
