package goopenmeteo

import (
	"encoding/json"
	"time"
)

type Forecast struct {
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	GenerationtimeMs float64 `json:"generationtime_ms"`
	UtcOffsetSeconds int     `json:"utc_offset_seconds"`
	Timezone         string  `json:"timezone"`
	TimezoneAbbr     string  `json:"timezone_abbreviation"`
	Elevation        float64 `json:"elevation"`
	HourlyUnits      Units   `json:"hourly_units"`
	Hourly           Hourly  `json:"hourly"`
}

type Units struct {
	Time          string `json:"time"`
	Temperature2m string `json:"temperature_2m"`
}

type Hourly struct {
	Time          []time.Time `json:"time"`
	Temperature2m []float64   `json:"temperature_2m"`
}

func (h *Hourly) UnmarshalJSON(data []byte) error {
	type Aux struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	}

	var aux Aux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	h.Time = make([]time.Time, len(aux.Time))
	for i, timeStr := range aux.Time {
		t, err := time.Parse("2006-01-02T15:04", timeStr)
		if err != nil {
			return err
		}
		h.Time[i] = t
	}

	h.Temperature2m = aux.Temperature2m
	return nil
}
