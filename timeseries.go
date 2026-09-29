package goopenmeteo

import (
	"encoding/json"
	"time"
)

type TimeseriesData struct {
	Time []time.Time                   `json:"time"`
	Data map[WeatherVariable][]float64 `json:"data"`
	// Times holds the series whose values are timestamps, such as daily
	// sunrise and sunset.
	Times map[WeatherVariable][]time.Time `json:"times,omitempty"`
}

func (t *TimeseriesData) UnmarshalJSON(data []byte) error {
	var rawData map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	var timeStrings []string
	if err := json.Unmarshal(rawData["time"], &timeStrings); err != nil {
		return err
	}

	var err error
	if t.Time, err = parseTimes(timeStrings); err != nil {
		return err
	}

	t.Data = make(map[WeatherVariable][]float64)

	for key, value := range rawData {
		if key == "time" {
			continue
		}

		// A null element decodes as 0 (floats) or the zero time (timestamps).
		var floatArr []float64
		if err := json.Unmarshal(value, &floatArr); err == nil {
			t.Data[key] = floatArr
			continue
		}

		var stringArr []string
		if err := json.Unmarshal(value, &stringArr); err != nil {
			return err
		}
		times, err := parseTimes(stringArr)
		if err != nil {
			return err
		}
		if t.Times == nil {
			t.Times = make(map[WeatherVariable][]time.Time)
		}
		t.Times[key] = times
	}

	return nil
}

// parseTimes parses Open-Meteo's ISO 8601 times ("2024-12-24T13:00") and dates
// ("2024-12-24"). An empty string (a null element) becomes the zero time.
func parseTimes(values []string) ([]time.Time, error) {
	times := make([]time.Time, len(values))
	for i, value := range values {
		if value == "" {
			continue
		}
		if len(value) == 10 { // Format: "2024-12-24"
			value = value + "T00:00"
		}
		var err error
		if times[i], err = time.Parse("2006-01-02T15:04", value); err != nil {
			return nil, err
		}
	}
	return times, nil
}
