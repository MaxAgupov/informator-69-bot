package wiki

import (
	"testing"
	"time"
)

func TestGetDayNoun(t *testing.T) {
	checkDays(t, "день", []int{1, 21, 31, 51, 61, 71, 81, 91, 101, 201, 301})
	checkDays(t, "дня", []int{2, 3, 4, 22, 23, 24, 102, 103, 104, 122, 123, 124})
	checkDays(t, "дней", []int{5, 6, 7, 8, 9, 10, 11, 12, 19, 20, 25, 100, 200, 300})
}

func checkDays(t *testing.T, noun string, vals []int) {
	for _,i := range vals {
		s := GetDayNoun(i)
		if s != noun {
			t.Error("Wrong format",  i, s, "( expected",i, noun, ")")
		}
	}
}

func TestGetCalendarInfoIssue(t *testing.T) {
	location, _ := time.LoadLocation(moscowLocation)
	var date = time.Date(2018, time.December, 7, 1,1,1,0, location)
	var calcInfo = GenerateCalendarStats(&date)
	var expectedInfo = "*Пятница, 7 декабря 2018 года*\n341-й день года. До конца года 24 дня\n"
	if expectedInfo != calcInfo {
		t.Error("Wrong line:", calcInfo, "(", expectedInfo, ")")
	}
}