package goopenmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

// fixtureServer serves a testdata file for one path and records the query it was asked.
func fixtureServer(t *testing.T, path, fixture string) (*httptest.Server, *url.Values) {
	t.Helper()
	body, err := os.ReadFile("testdata/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	var query url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.Error(w, `{"error":true,"reason":"wrong path `+r.URL.Path+`"}`, http.StatusBadRequest)
			return
		}
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}))
	t.Cleanup(srv.Close)
	return srv, &query
}

func newTestClient(srv *httptest.Server) *OpenMeteo {
	o := NewOpenMeteo()
	o.BaseURL = srv.URL
	o.AirQualityURL = srv.URL
	return o
}

var zagreb = time.FixedZone("GMT+2", 7200)

func TestForecastOptionsQueryUsesOpenMeteoParameterNames(t *testing.T) {
	opts := ForecastOptions{
		Latitude:          45.5,
		Longitude:         17.2,
		Minutely15:        WeatherVariables{Temperature2M},
		TemperatureUnit:   TemperatureUnitFahrenheit,
		WindSpeedUnit:     WindSpeedUnitMS,
		PrecipitationUnit: PrecipitationUnitInch,
		PastDays:          1,
		ForecastDays:      3,
		ForecastHours:     12,
		PastHours:         2,
		PastMinutely15:    4,
		StartDate:         "2026-09-01",
		EndDate:           "2026-09-02",
		StartHour:         "2026-09-01T00:00",
		EndHour:           "2026-09-01T12:00",
		StartMinutely15:   "2026-09-01T00:00",
		EndMinutely15:     "2026-09-01T01:00",
		CellSelection:     CellSelectionLand,
	}
	q, err := url.ParseQuery(opts.ToQuery())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"minutely_15":        "temperature_2m",
		"temperature_unit":   "fahrenheit",
		"wind_speed_unit":    "ms",
		"precipitation_unit": "inch",
		"past_days":          "1",
		"forecast_days":      "3",
		"forecast_hours":     "12",
		"past_hours":         "2",
		"past_minutely_15":   "4",
		"start_date":         "2026-09-01",
		"end_date":           "2026-09-02",
		"start_hour":         "2026-09-01T00:00",
		"end_hour":           "2026-09-01T12:00",
		"start_minutely_15":  "2026-09-01T00:00",
		"end_minutely_15":    "2026-09-01T01:00",
		"cell_selection":     "land",
	}
	for key, value := range want {
		if got := q.Get(key); got != value {
			t.Errorf("query %s = %q, want %q (full query %s)", key, got, value, q.Encode())
		}
	}
}

func TestForecastContextDecodesResponseMetadata(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")

	f, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{Latitude: 45.59, Longitude: 17.22})
	if err != nil {
		t.Fatal(err)
	}
	if f.UTCOffsetSeconds != 7200 {
		t.Errorf("UTCOffsetSeconds = %d, want 7200", f.UTCOffsetSeconds)
	}
	if f.TimezoneAbbreviation != "GMT+2" {
		t.Errorf("TimezoneAbbreviation = %q, want GMT+2", f.TimezoneAbbreviation)
	}
	if f.GenerationTimeMS <= 0 {
		t.Errorf("GenerationTimeMS = %v, want > 0", f.GenerationTimeMS)
	}
	if f.CurrentUnits[Temperature2M] != "°C" || f.HourlyUnits[PrecipitationProbability] != "%" || f.DailyUnits[Temperature2MMax] != "°C" {
		t.Errorf("units not decoded: current %v hourly %v daily %v", f.CurrentUnits, f.HourlyUnits, f.DailyUnits)
	}
	if f.Current.Data[Temperature2M] != 12.5 || f.Current.Data[IsDay] != 0 {
		t.Errorf("current data = %v", f.Current.Data)
	}
}

func TestForecastTimesAreInstantsInTheResponseZone(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")

	f, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if want := time.Date(2026, 9, 29, 1, 0, 0, 0, zagreb); !f.Current.Time.Equal(want) {
		t.Errorf("Current.Time = %v, want %v", f.Current.Time, want)
	}
	if want := time.Date(2026, 9, 29, 0, 0, 0, 0, zagreb); !f.Hourly.Time[0].Equal(want) {
		t.Errorf("Hourly.Time[0] = %v, want %v", f.Hourly.Time[0], want)
	}
	if want := time.Date(2026, 9, 30, 0, 0, 0, 0, zagreb); !f.Daily.Time[1].Equal(want) {
		t.Errorf("Daily.Time[1] = %v, want %v", f.Daily.Time[1], want)
	}
	if _, offset := f.Hourly.Time[5].Zone(); offset != 7200 {
		t.Errorf("hourly offset = %d, want 7200", offset)
	}
}

func TestForecastWithoutTimezoneKeepsUTC(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"utc_offset_seconds":0,"timezone":"GMT","timezone_abbreviation":"GMT",
			"current":{"time":"2026-09-29T01:00","interval":900,"temperature_2m":12.5}}`))
	}))
	defer srv.Close()

	f, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if want := time.Date(2026, 9, 29, 1, 0, 0, 0, time.UTC); f.Current.Time != want {
		t.Errorf("Current.Time = %v, want %v", f.Current.Time, want)
	}
}

func TestDailyStringSeriesDecodeAsTimes(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")

	f, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{})
	if err != nil {
		t.Fatal(err)
	}
	sunrise := f.Daily.Times[Sunrise]
	if len(sunrise) != 2 {
		t.Fatalf("Daily.Times[sunrise] = %v, want 2 entries", sunrise)
	}
	if want := time.Date(2026, 9, 29, 6, 46, 0, 0, zagreb); !sunrise[0].Equal(want) {
		t.Errorf("sunrise[0] = %v, want %v", sunrise[0], want)
	}
	if want := time.Date(2026, 9, 30, 18, 33, 0, 0, zagreb); !f.Daily.Times[Sunset][1].Equal(want) {
		t.Errorf("sunset[1] = %v, want %v", f.Daily.Times[Sunset][1], want)
	}
	if got := f.Daily.Data[Temperature2MMax]; len(got) != 2 || got[1] != 25.6 {
		t.Errorf("Daily.Data[temperature_2m_max] = %v", got)
	}
	if _, isFloat := f.Daily.Data[Sunrise]; isFloat {
		t.Error("sunrise must not appear in Data")
	}
}

func TestTimeseriesToleratesNullElements(t *testing.T) {
	var ts TimeseriesData
	err := json.Unmarshal([]byte(`{"time":["2026-09-29T00:00","2026-09-29T01:00"],
		"temperature_2m":[1.5,null],"sunrise":["2026-09-29T06:46",null]}`), &ts)
	if err != nil {
		t.Fatal(err)
	}
	if got := ts.Data[Temperature2M]; len(got) != 2 || got[0] != 1.5 {
		t.Errorf("Data[temperature_2m] = %v", got)
	}
	if got := ts.Times[Sunrise]; len(got) != 2 || !got[1].IsZero() {
		t.Errorf("Times[sunrise] = %v, want a zero time for null", got)
	}
}

func TestForecastSendsOptionsToForecastEndpoint(t *testing.T) {
	srv, query := fixtureServer(t, "/forecast", "forecast.json")

	_, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{
		Latitude:     45.59,
		Longitude:    17.22,
		Current:      WeatherVariables{Temperature2M, IsDay},
		Timezone:     "Europe/Zagreb",
		ForecastDays: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if query.Get("current") != "temperature_2m,is_day" || query.Get("forecast_days") != "2" || query.Get("timezone") != "Europe/Zagreb" {
		t.Errorf("query = %v", *query)
	}
}

func TestForecastWithoutContextStillWorks(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")

	f, err := newTestClient(srv).Forecast(ForecastOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if f.Current.Data[Temperature2M] != 12.5 {
		t.Errorf("current = %v", f.Current.Data)
	}
}

func TestZeroValueClientUsesDefaults(t *testing.T) {
	o := &OpenMeteo{}
	if got := o.forecastURL(); got != BaseURL {
		t.Errorf("forecast base = %q, want %q", got, BaseURL)
	}
	if got := o.airQualityURL(); got != AirQualityBaseURL {
		t.Errorf("air quality base = %q, want %q", got, AirQualityBaseURL)
	}
	if o.client() == nil {
		t.Error("client() = nil")
	}
}

func TestForecastContextHonoursCancellation(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := newTestClient(srv).ForecastContext(ctx, ForecastOptions{})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestForecastReturnsTheAPIReasonOnBadRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true,"reason":"Forecast days is invalid. Allowed range 0 to 16."}`))
	}))
	defer srv.Close()

	_, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{ForecastDays: 99})
	if err == nil || !strings.Contains(err.Error(), "Forecast days is invalid") {
		t.Errorf("err = %v, want the API reason", err)
	}
}

func TestForecastReportsUnexpectedStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	_, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{})
	if err == nil || !strings.Contains(err.Error(), "503") {
		t.Errorf("err = %v, want it to name status 503", err)
	}
}

func TestForecastMarshalsCamelCaseKeys(t *testing.T) {
	srv, _ := fixtureServer(t, "/forecast", "forecast.json")

	f, err := newTestClient(srv).ForecastContext(context.Background(), ForecastOptions{})
	if err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{`"utcOffsetSeconds":7200`, `"timezoneAbbreviation":"GMT+2"`, `"hourlyUnits"`} {
		if !strings.Contains(string(out), key) {
			t.Errorf("marshalled forecast lacks %s: %s", key, out[:200])
		}
	}
}

func TestAirQualityContextDecodesResponse(t *testing.T) {
	srv, query := fixtureServer(t, "/air-quality", "airquality.json")

	aq, err := newTestClient(srv).AirQualityContext(context.Background(), AirQualityOptions{
		Latitude:     45.59,
		Longitude:    17.22,
		Current:      AirQualityVariables{EuropeanAQI, RagweedPollen},
		Hourly:       AirQualityVariables{BirchPollen, EuropeanAQI},
		Timezone:     "Europe/Zagreb",
		ForecastDays: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if query.Get("current") != "european_aqi,ragweed_pollen" || query.Get("forecast_days") != "1" {
		t.Errorf("query = %v", *query)
	}
	if aq.Current.Data[EuropeanAQI] != 23 || aq.Current.Data[RagweedPollen] != 1.4 {
		t.Errorf("current = %v", aq.Current.Data)
	}
	if got := aq.Hourly.Data[EuropeanAQI]; len(got) != 24 || got[16] != 31 {
		t.Errorf("hourly european_aqi = %v", got)
	}
	if want := time.Date(2026, 9, 29, 1, 0, 0, 0, zagreb); !aq.Current.Time.Equal(want) {
		t.Errorf("Current.Time = %v, want %v", aq.Current.Time, want)
	}
	if aq.HourlyUnits[BirchPollen] != "grains/m³" {
		t.Errorf("hourly units = %v", aq.HourlyUnits)
	}
}
