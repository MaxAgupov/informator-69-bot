package weather

import (
	"github.com/briandowns/openweathermap"
	"log"
	"testing"
)

func TestCurrentWeather(t *testing.T) {
	w := GetCurrentWeather("Voronezh")
	log.Println(w)
}

func TestForecast(t *testing.T) {

	w, err := openweathermap.NewForecast( "5", "C", "RU", apiKey)
	if err != nil {
		log.Fatal(err)
	}
	err = w.DailyByName("Voronezh,RU", 10)
	if err != nil {
		log.Fatal(err)
	}
}