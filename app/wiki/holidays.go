package wiki

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type DayHolidays struct {
	Month  string `json:"month"`
	Day    string `json:"day"`
	Report Report `json:"report"`
}

type MonthHolidays map[int]*DayHolidays
type Holidays map[time.Month]*MonthHolidays

func LoadHolidays(holidaysFileName string) Holidays {
	var holidays = Holidays{}

	file, err := os.OpenFile(holidaysFileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Print(err)
		}
	}()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&holidays); err != nil {
		log.Print(err)
	}

	return holidays
}
