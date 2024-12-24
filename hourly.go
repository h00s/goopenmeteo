package goopenmeteo

import (
	"encoding/json"
	"time"
)

type Hourly struct {
	Time []time.Time          `json:"time"`
	Data map[string][]float64 `json:"-"`
}

func (h *Hourly) UnmarshalJSON(data []byte) error {
	var rawData map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	var timeStrings []string
	if err := json.Unmarshal(rawData["time"], &timeStrings); err != nil {
		return err
	}

	h.Time = make([]time.Time, len(timeStrings))
	for i, timeStr := range timeStrings {
		t, err := time.Parse("2006-01-02T15:04", timeStr)
		if err != nil {
			return err
		}
		h.Time[i] = t
	}

	h.Data = make(map[string][]float64)

	for key, value := range rawData {
		if key == "time" {
			continue
		}

		var floatArr []float64
		if err := json.Unmarshal(value, &floatArr); err != nil {
			return err
		}
		h.Data[key] = floatArr
	}

	return nil
}
