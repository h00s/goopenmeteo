package goopenmeteo

import (
	"encoding/json"
	"time"
)

type Current struct {
	Time     time.Time                   `json:"time"`
	Interval int                         `json:"interval"`
	Data     map[WeatherVariable]float64 `json:"data"`
}

func (c *Current) UnmarshalJSON(data []byte) error {
	var rawData map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	var timeString string
	if err := json.Unmarshal(rawData["time"], &timeString); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04", timeString)
	if err != nil {
		return err
	}
	c.Time = t

	if err := json.Unmarshal(rawData["interval"], &c.Interval); err != nil {
		return err
	}

	c.Data = make(map[WeatherVariable]float64)

	for key, value := range rawData {
		if key == "time" || key == "interval" {
			continue
		}

		var floatVal float64
		if err := json.Unmarshal(value, &floatVal); err != nil {
			return err
		}
		c.Data[key] = floatVal
	}

	return nil
}
