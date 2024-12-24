package goopenmeteo

type WeatherVariable = string
type WeatherVariables = []string

const (
	// Temperature and Comfort
	Temperature2M       WeatherVariable = "temperature_2m"
	RelativeHumidity2M  WeatherVariable = "relative_humidity_2m"
	DewPoint2M          WeatherVariable = "dew_point_2m"
	ApparentTemperature WeatherVariable = "apparent_temperature"

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

	// Solar Radiation
	ShortwaveRadiation     WeatherVariable = "shortwave_radiation"
	DirectRadiation        WeatherVariable = "direct_radiation"
	DirectNormalIrradiance WeatherVariable = "direct_normal_irradiance"
	DiffuseRadiation       WeatherVariable = "diffuse_radiation"
	GlobalTiltedIrradiance WeatherVariable = "global_tilted_irradiance"

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
