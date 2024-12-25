package goopenmeteo

import (
	"encoding/json"
	"fmt"
)

const BaseURL = "https://api.open-meteo.com/v1"

type OpenMeteo struct {
	APIKey string
}

func NewOpenMeteo(APIKey ...string) *OpenMeteo {
	if len(APIKey) > 0 {
		return &OpenMeteo{APIKey: APIKey[0]}
	}
	return &OpenMeteo{}
}

func (o *OpenMeteo) Forecast(options ForecastOptions) (*Forecast, error) {
	url := fmt.Sprintf("%s/forecast?%s", BaseURL, options.ToQuery())
	if o.APIKey != "" {
		url = fmt.Sprintf("%s&apikey=%s", url, o.APIKey)
	}
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
