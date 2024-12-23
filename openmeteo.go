package goopenmeteo

import (
	"encoding/json"
	"fmt"
)

const BaseURL = "https://api.open-meteo.com/v1"

type OpenMeteo struct {
	APIKey string
}

func NewOpenMeteo(APIKey string) *OpenMeteo {
	return &OpenMeteo{APIKey: APIKey}
}

// Forecast returns the forecast for the given latitude and longitude.
func (o *OpenMeteo) Forecast(latitude, longitude float64) (*Forecast, error) {
	url := fmt.Sprintf("%s/forecast?latitude=%f&longitude=%f&hourly=temperature_2m", BaseURL, latitude, longitude)
	response, err := httpGet(url)
	if err != nil {
		return nil, err
	}

	var forecast Forecast
	if err := json.Unmarshal(response, &forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}
