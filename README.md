# Golang Open-Meteo

Open-Meteo is an open-source weather API and offers free access for non-commercial use. No API key required.

```bash
go get github.com/h00s/goopenmeteo
```

## Forecast

```go
om := goopenmeteo.NewOpenMeteo() // or NewOpenMeteo(apiKey)

forecast, err := om.ForecastContext(ctx, goopenmeteo.ForecastOptions{
	Latitude:     45.5936,
	Longitude:    17.2251,
	Current:      goopenmeteo.WeatherVariables{goopenmeteo.Temperature2M, goopenmeteo.WeatherCode},
	Hourly:       goopenmeteo.WeatherVariables{goopenmeteo.Temperature2M, goopenmeteo.PrecipitationProbability},
	Daily:        goopenmeteo.WeatherVariables{goopenmeteo.Temperature2MMax, goopenmeteo.Sunrise, goopenmeteo.Sunset},
	Timezone:     "Europe/Zagreb",
	ForecastDays: 3,
})

forecast.Current.Data[goopenmeteo.Temperature2M]   // float64
forecast.Hourly.Time                               // []time.Time
forecast.Hourly.Data[goopenmeteo.Temperature2M]    // []float64, aligned with Hourly.Time
forecast.Daily.Times[goopenmeteo.Sunrise]          // []time.Time: series whose values are timestamps
```

`Forecast(options)` is the same call with a background context.

- **Times are instants.** Open-Meteo sends local wall-clock times without an offset; they are placed in the response's zone (`time.FixedZone(TimezoneAbbreviation, UTCOffsetSeconds)`). Without a `Timezone` option the API answers in GMT and times are UTC.
- **Numeric series** are in `Data`, **timestamp series** (`sunrise`, `sunset`) in `Times`. A `null` element decodes as `0` or the zero time.
- **Output JSON** uses camelCase keys (`utcOffsetSeconds`, `hourlyUnits`, …), so a `Forecast` can be passed straight to a JavaScript client.

## Air quality and pollen

```go
aq, err := om.AirQualityContext(ctx, goopenmeteo.AirQualityOptions{
	Latitude:     45.5936,
	Longitude:    17.2251,
	Current:      goopenmeteo.AirQualityVariables{goopenmeteo.EuropeanAQI, goopenmeteo.RagweedPollen},
	Hourly:       goopenmeteo.AirQualityVariables{goopenmeteo.BirchPollen, goopenmeteo.GrassPollen},
	Timezone:     "Europe/Zagreb",
	ForecastDays: 1,
})
```

Pollen is available for Europe only, during the pollen season.

## Configuration

`OpenMeteo` fields, all optional (the zero value works):

| Field | Default |
|---|---|
| `APIKey` | none (free API) |
| `BaseURL` | `https://api.open-meteo.com/v1` |
| `AirQualityURL` | `https://air-quality-api.open-meteo.com/v1` |
| `HTTPClient` | an `*http.Client` with a 10 s timeout |

Point `BaseURL` / `AirQualityURL` at an `httptest.Server` in tests.

## Errors

A `400` returns Open-Meteo's `reason` (`open-meteo: Forecast days is invalid…`); any other non-200 status names the status code. Context cancellation is returned as-is (`errors.Is(err, context.Canceled)`).

## Tests

```bash
go test ./...                       # offline, against recorded fixtures
go test -tags live -run Live ./...  # against the real API
```
