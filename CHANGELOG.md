# Changelog

## v1.1.0

### Fixed
- **Query parameters use Open-Meteo's names again** (`forecast_days`, `temperature_unit`, `minutely_15`, `cell_selection`, …). v1.0.1 sent camelCase names, which the API silently ignored.
- **Response metadata decodes again.** `generationtime_ms`, `utc_offset_seconds`, `timezone_abbreviation` and the `*_units` maps were never filled in v1.0.1. Output JSON keeps the camelCase keys introduced in v1.0.1.
- **Timestamp series decode.** Daily `sunrise` and `sunset` made the whole response fail; they now land in the new `TimeseriesData.Times`.
- **A 400 returns the API's reason** instead of an empty error; other statuses name the code.

### Changed
- **Times are placed in the response's zone** when a `Timezone` is requested, so they are real instants. Previously local wall-clock times were labelled UTC. Responses in GMT (no `Timezone` option) are unchanged.

### Added
- `ForecastContext(ctx, options)`; `Forecast` delegates with a background context.
- `AirQualityContext(ctx, options)` with `AirQualityOptions`, `AirQuality` and air quality / pollen variable constants.
- `OpenMeteo.BaseURL`, `OpenMeteo.AirQualityURL` and `OpenMeteo.HTTPClient`; the zero value `&OpenMeteo{}` is usable.
- Tests against recorded fixtures, plus a live test behind the `live` build tag.

## v1.0.1

- JSON output keys changed to camelCase.

## v1.0.0

- Initial release.
