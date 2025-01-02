package goopenmeteo

import (
	"encoding/json"
	"time"
)

type TimeseriesData struct {
	Time []time.Time                   `json:"time"`
	Data map[WeatherVariable][]float64 `json:"data"`
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

	t.Time = make([]time.Time, len(timeStrings))
	for i, timeStr := range timeStrings {
		var err error
		if len(timeStr) == 10 { // Format: "2024-12-24"
			timeStr = timeStr + "T00:00"
		}
		t.Time[i], err = time.Parse("2006-01-02T15:04", timeStr)
		if err != nil {
			return err
		}
	}

	t.Data = make(map[WeatherVariable][]float64)

	for key, value := range rawData {
		if key == "time" {
			continue
		}

		var floatArr []float64
		if err := json.Unmarshal(value, &floatArr); err != nil {
			return err
		}
		t.Data[key] = floatArr
	}

	return nil
}
