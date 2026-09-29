package goopenmeteo

import (
	"encoding/json"
	"time"
)

// responseWire is the shape Open-Meteo sends (snake_case keys). The exported
// types keep their camelCase JSON tags for output, so they decode through this.
type responseWire struct {
	Latitude             float64           `json:"latitude"`
	Longitude            float64           `json:"longitude"`
	GenerationTimeMS     float64           `json:"generationtime_ms"`
	UTCOffsetSeconds     int               `json:"utc_offset_seconds"`
	Timezone             string            `json:"timezone"`
	TimezoneAbbreviation string            `json:"timezone_abbreviation"`
	Elevation            float64           `json:"elevation"`
	CurrentUnits         map[string]string `json:"current_units"`
	Current              Current           `json:"current"`
	Minutely15Units      map[string]string `json:"minutely_15_units"`
	Minutely15           TimeseriesData    `json:"minutely_15"`
	HourlyUnits          map[string]string `json:"hourly_units"`
	Hourly               TimeseriesData    `json:"hourly"`
	DailyUnits           map[string]string `json:"daily_units"`
	Daily                TimeseriesData    `json:"daily"`
}

// UnmarshalJSON decodes an Open-Meteo forecast response. Times in the response
// are local wall-clock times without an offset; they are placed in the
// response's zone so they are real instants.
func (f *Forecast) UnmarshalJSON(data []byte) error {
	var w responseWire
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	loc := w.location()
	*f = Forecast{
		Latitude:             w.Latitude,
		Longitude:            w.Longitude,
		GenerationTimeMS:     w.GenerationTimeMS,
		UTCOffsetSeconds:     w.UTCOffsetSeconds,
		Timezone:             w.Timezone,
		TimezoneAbbreviation: w.TimezoneAbbreviation,
		Elevation:            w.Elevation,
		CurrentUnits:         w.CurrentUnits,
		Current:              w.Current.in(loc),
		Minutely15Units:      w.Minutely15Units,
		Minutely15:           w.Minutely15.in(loc),
		HourlyUnits:          w.HourlyUnits,
		Hourly:               w.Hourly.in(loc),
		DailyUnits:           w.DailyUnits,
		Daily:                w.Daily.in(loc),
	}
	return nil
}

// location is nil for a UTC response, which keeps the parsed times as they are.
func (w *responseWire) location() *time.Location {
	if w.UTCOffsetSeconds == 0 {
		return nil
	}
	return time.FixedZone(w.TimezoneAbbreviation, w.UTCOffsetSeconds)
}

// inZone reinterprets a wall-clock time parsed as UTC in loc.
func inZone(t time.Time, loc *time.Location) time.Time {
	if loc == nil || t.IsZero() {
		return t
	}
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

func (c Current) in(loc *time.Location) Current {
	c.Time = inZone(c.Time, loc)
	return c
}

func (t TimeseriesData) in(loc *time.Location) TimeseriesData {
	if loc == nil {
		return t
	}
	for i := range t.Time {
		t.Time[i] = inZone(t.Time[i], loc)
	}
	for _, series := range t.Times {
		for i := range series {
			series[i] = inZone(series[i], loc)
		}
	}
	return t
}
