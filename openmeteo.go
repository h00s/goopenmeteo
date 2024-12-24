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

func (o *OpenMeteo) Forecast(options ForecastOptions) (*Forecast, error) {
	url := fmt.Sprintf("%s/forecast?%s", BaseURL, options.Query())
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
