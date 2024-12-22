package goopenmeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var forecast Forecast
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}
