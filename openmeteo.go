package goopenmeteo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// BaseURL is the default root of the forecast API.
const BaseURL = "https://api.open-meteo.com/v1"

// AirQualityBaseURL is the default root of the air quality API.
const AirQualityBaseURL = "https://air-quality-api.open-meteo.com/v1"

const defaultTimeout = 10 * time.Second

var defaultClient = &http.Client{Timeout: defaultTimeout}

// OpenMeteo is an Open-Meteo API client. The zero value is usable and talks to
// the public API.
type OpenMeteo struct {
	APIKey string
	// BaseURL overrides the forecast API root, BaseURL by default.
	BaseURL string
	// AirQualityURL overrides the air quality API root, AirQualityBaseURL by default.
	AirQualityURL string
	// HTTPClient sends the requests; a client with a 10 s timeout by default.
	HTTPClient *http.Client
}

func NewOpenMeteo(APIKey ...string) *OpenMeteo {
	o := &OpenMeteo{
		BaseURL:       BaseURL,
		AirQualityURL: AirQualityBaseURL,
		HTTPClient:    &http.Client{Timeout: defaultTimeout},
	}
	if len(APIKey) > 0 {
		o.APIKey = APIKey[0]
	}
	return o
}

// Forecast is ForecastContext with a background context.
func (o *OpenMeteo) Forecast(options ForecastOptions) (*Forecast, error) {
	return o.ForecastContext(context.Background(), options)
}

// ForecastContext requests a weather forecast; ctx cancels the request.
func (o *OpenMeteo) ForecastContext(ctx context.Context, options ForecastOptions) (*Forecast, error) {
	var forecast Forecast
	if err := o.get(ctx, o.forecastURL()+"/forecast", options.ToQuery(), &forecast); err != nil {
		return nil, err
	}
	return &forecast, nil
}

// AirQualityContext requests air quality and pollen data; ctx cancels the request.
func (o *OpenMeteo) AirQualityContext(ctx context.Context, options AirQualityOptions) (*AirQuality, error) {
	var airQuality AirQuality
	if err := o.get(ctx, o.airQualityURL()+"/air-quality", options.ToQuery(), &airQuality); err != nil {
		return nil, err
	}
	return &airQuality, nil
}

func (o *OpenMeteo) get(ctx context.Context, endpoint, query string, v any) error {
	if o.APIKey != "" {
		query += "&apikey=" + url.QueryEscape(o.APIKey)
	}
	response, err := httpGet(ctx, o.client(), endpoint+"?"+query)
	if err != nil {
		return err
	}
	return json.Unmarshal(response, v)
}

func (o *OpenMeteo) forecastURL() string {
	if o.BaseURL != "" {
		return o.BaseURL
	}
	return BaseURL
}

func (o *OpenMeteo) airQualityURL() string {
	if o.AirQualityURL != "" {
		return o.AirQualityURL
	}
	return AirQualityBaseURL
}

func (o *OpenMeteo) client() *http.Client {
	if o.HTTPClient != nil {
		return o.HTTPClient
	}
	return defaultClient
}
