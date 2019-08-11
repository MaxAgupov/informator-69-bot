package wiki

import (
	"testing"
	"time"
)

var testFileName = "holidays_test.json"

func TestLoading(t *testing.T) {
	holi := LoadHolidays(testFileName)
	jan := holi[time.January]
	if jan == nil {
		t.Error("Can't load January")
	}
	feb := holi[time.February]
	if feb != nil {
		t.Error("Load February")
	}

	first := (*jan)[1]
	if first == nil {
		t.Error("Can't load day")
	}
}
