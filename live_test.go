//go:build live

package goopenmeteo

import (
	"context"
	"testing"
	"time"
)

// Run with: go test -tags live -run Live
func TestLiveForecastAndAirQuality(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	o := NewOpenMeteo()

	f, err := o.ForecastContext(ctx, ForecastOptions{
		Latitude:     45.5936,
		Longitude:    17.2251,
		Current:      WeatherVariables{Temperature2M, WeatherCode, IsDay},
		Hourly:       WeatherVariables{Temperature2M, PrecipitationProbability},
		Daily:        WeatherVariables{Temperature2MMax, Sunrise, Sunset},
		Timezone:     "Europe/Zagreb",
		ForecastDays: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	if f.UTCOffsetSeconds == 0 || len(f.Daily.Time) != 3 || len(f.Daily.Times[Sunrise]) != 3 || len(f.Hourly.Time) != 72 {
		t.Errorf("forecast: offset %d, daily %d, sunrise %d, hourly %d", f.UTCOffsetSeconds, len(f.Daily.Time), len(f.Daily.Times[Sunrise]), len(f.Hourly.Time))
	}

	aq, err := o.AirQualityContext(ctx, AirQualityOptions{
		Latitude:     45.5936,
		Longitude:    17.2251,
		Current:      AirQualityVariables{EuropeanAQI, RagweedPollen, BirchPollen},
		Timezone:     "Europe/Zagreb",
		ForecastDays: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := aq.Current.Data[EuropeanAQI]; !ok {
		t.Errorf("air quality current = %v", aq.Current.Data)
	}
}
