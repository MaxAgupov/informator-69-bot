package weather

import (
	"github.com/briandowns/openweathermap"
	"log"
)

var apiKey = "912a4fd282aacbdd2d9d2c6da9c85d2f"

func GetCurrentWeather(city string) *openweathermap.CurrentWeatherData {
	w, err := openweathermap.NewCurrent("C", "EN", apiKey)
	if err != nil {
		log.Fatal(err)
	}
	err = w.CurrentByName(city)
	if err != nil {
		log.Fatal(err)
	}
	return w
}

func GetForecast(city string) *openweathermap.ForecastWeatherData {
	w, err := openweathermap.NewForecast("5","C", "EN", apiKey)
	if err != nil {
		log.Fatal(err)
	}
	w.DailyByName(city, 10)
	if err != nil {
		log.Fatal(err)
	}
	return w
}


