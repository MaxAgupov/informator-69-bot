package wiki

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var monthsGenitive = [...]string{
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

var weekDays = [...]string{
	"воскресенье",
	"понедельник",
	"вторник",
	"среда",
	"четверг",
	"пятница",
	"суббота",
}

const holidaysHeader = "Праздники и памятные дни"
const intHolidaysSubheader = "Международные"
const locHolidaysSubheader = "Национальные"
const rlgHolidaysSubheader = "Религиозные"
const profHolidaysSubheader = "Профессиональные"
const nameDaysSubheader = "Именины"

const moscowLocation = "Europe/Moscow"

var reportCache = ReportCache{}

type Report struct {
	stats        string
	common       []string
	holidaysInt  []string
	holidaysLoc  []string
	holidaysProf []string
	holidaysRlg  ReligiousHolidays
	nameDays     []string
	sections     map[string][]*Section
}

type ReligiousHolidays struct {
	orthodox    []string
	catholicism []string
	others      []string
}

func (holidays *ReligiousHolidays) Empty() bool {
	return len(holidays.orthodox) == 0 && len(holidays.catholicism) == 0 && len(holidays.others) == 0
}

func (holidays *ReligiousHolidays) AppendString(formatted *string) {
	if len(holidays.orthodox) > 0 {
		for _, line := range holidays.orthodox {
			*formatted += "- " + line + " (правосл.)\n"
		}
	}
	if len(holidays.catholicism) > 0 {
		for _, line := range holidays.catholicism {
			*formatted += "- " + line + " (катол.)\n"
		}
	}
	if len(holidays.others) > 0 {
		for _, line := range holidays.others {
			*formatted += "- " + line + "\n"
		}
	}
}

type Section struct {
	header  string
	content []string
}

func (report *Report) String() string {
	formattedStr := ""
	if report.stats != "" {
		formattedStr += report.stats + "\n"
	}

	if len(report.holidaysInt) > 0 || len(report.holidaysLoc) > 0 || len(report.holidaysProf) > 0 || !report.holidaysRlg.Empty() {
		formattedStr += "*" + holidaysHeader + "*\n"
		if len(report.holidaysInt) > 0 {
			formattedStr += "\n_" + intHolidaysSubheader + "_\n"
			for _, line := range report.holidaysInt {
				formattedStr += "- " + line + "\n"
			}
		}
		if len(report.holidaysLoc) > 0 {
			formattedStr += "\n_" + locHolidaysSubheader + "_\n"
			for _, line := range report.holidaysLoc {
				formattedStr += "- " + line + "\n"
			}
		}
		if len(report.holidaysProf) > 0 {
			formattedStr += "\n_" + profHolidaysSubheader + "_\n"
			for _, line := range report.holidaysProf {
				formattedStr += "- " + line + "\n"
			}
		}
		if !report.holidaysRlg.Empty() {
			formattedStr += "\n_" + rlgHolidaysSubheader + "_\n"
			report.holidaysRlg.AppendString(&formattedStr)
		}
	}

	if len(report.nameDays) > 0 {
		formattedStr += "\n_" + nameDaysSubheader + "_\n"
		for _, line := range report.nameDays {
			formattedStr += "- " + line + "\n"
		}
	}
	return formattedStr
}

func (report *Report) Events() string {
	formattedStr := ""
	if report.stats != "" {
		formattedStr += report.stats + "\n"
	}
	for k, v := range report.sections {
		if k == "События" || k == "Приметы" {
			formattedStr += "*" + k + "*\n"
			for _, sect := range v {
				formattedStr += "\n_" + sect.header + "_" + "\n"
				formattedStr += strings.Join(sect.content, "\n") + "\n"
			}
			formattedStr += "\n"
		}
	}
	return formattedStr
}

func (report *Report) setCalendarInfo(day *time.Time) {
	report.stats = GenerateCalendarStats(day)
}

type Response struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

type Query struct {
	Pages map[string]Pages `json:"pages"`
}

type Pages struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
	PageId  uint64 `json:"pageid"`
	NS      uint64 `json:"ns"`
}

func getWikiReport(reportDay *time.Time) string {
	wikiRequest := "https://ru.wikipedia.org/w/api.php?action=query&format=json&&prop=extracts&exlimit=1&explaintext"
	data := getDateString(reportDay)
	wikiRequest += "&titles=" + url.QueryEscape(data)

	log.Print(wikiRequest)
	if response, err := http.Get(wikiRequest); err != nil {
		log.Print("Wikipedia is not respond")
	} else {
		defer func() {
			if err := response.Body.Close(); err != nil {
				log.Print(err)
			}
		}()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
			return ""
		}
		var wr Response
		if err := json.Unmarshal(contents, &wr); err != nil {
			log.Print("Error", err)
			return ""
		}

		if l := len(wr.Query.Pages); l == 0 || l > 1 {
			log.Print("There must be only one page - ", l)
			return ""
		}
		var content string
		for _, v := range wr.Query.Pages {
			content = v.Extract
		}
		return content
	}
	return ""
}

func GetTodaysReport() string {
	location, _ := time.LoadLocation(moscowLocation)
	log.Print(location)
	now := time.Now().In(location)
	report := reportCache.getCachedReport(&now)
	return report.String()
}

func getDateString(day *time.Time) string {
	_, month, dayNum := day.Date()
	return strconv.Itoa(dayNum) + " " + monthsGenitive[month-1]
}

func getFullDateString(day *time.Time) string {
	year, month, dayNum := day.Date()
	weekDay := strings.Title(getWeekDateString(day))
	return "*" + weekDay + ", " + strconv.Itoa(dayNum) + " " + monthsGenitive[month-1] + " " + strconv.Itoa(year) + " года" + "*"
}

func getWeekDateString(day *time.Time) string {
	weekday := int(day.Weekday())
	return weekDays[weekday]
}

func GetDayNoun(day int) string {
	rest := day % 10
	if (day > 10) && (day < 20) {
		// для второго десятка - всегда третья форма
		return "дней"
	} else if rest == 1 {
		return "день"
	} else if rest > 1 && rest < 5 {
		return "дня"
	} else {
		return "дней"
	}
}

func GenerateCalendarStats(reportDay *time.Time) string {
	firstLine := getFullDateString(reportDay)

	year := time.Date(reportDay.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	infoDay := reportDay.YearDay()
	full_days := year.YearDay()

	rest := full_days - infoDay
	secondLine := ""
	if rest > 0 {
		secondLine = strconv.Itoa(infoDay) + "-й день года. До конца года " + strconv.Itoa(rest) + " " + GetDayNoun(rest)
	} else {
		secondLine = "Завтра уже Новый Год!"
	}

	return firstLine + "\n" + secondLine + "\n"
}

type ReportCache struct {
	sync.Mutex
	year   int
	month  time.Month
	day    int
	report *Report
}

func (cache *ReportCache) getCachedReport(date *time.Time) Report {
	cache.Lock()
	defer cache.Unlock()
	year, month, day := date.Date()
	if year == cache.year && month == cache.month && day == cache.day {
		return *cache.report
	}
	fullReport := getWikiReport(date)
	report, err := Parse(fullReport)
	if err != nil {
		log.Print("Error:", err)
		return Report{}
	}
	report.setCalendarInfo(date)
	cache.report = &report
	cache.year = year
	cache.month = month
	cache.day = day

	return *cache.report
}
