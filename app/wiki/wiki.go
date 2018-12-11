package wiki

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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

type Report struct {
	calendarInfo string
	common       []string
	sections     map[string][]*Section
}

type Section struct {
	header  string
	content []string
}

func (report *Report) String() string {
	formattedStr := ""
	if report.calendarInfo != "" {
		formattedStr += report.calendarInfo + "\n"
	}
	for k, v := range report.sections {
		//log.Print(k)
		if k == "Праздники и памятные дни" {
			formattedStr += "*" + k + "*\n"
			for _, sect := range v {
				if len(sect.content) == 0 {
					continue
				}
				formattedStr += "\n_" + sect.header + "_" + "\n"
				formattedStr += strings.Join(sect.content, "\n") + "\n"
			}
			formattedStr += "\n"
		}
	}
	return formattedStr
}

func (report *Report) Events() string {
	formattedStr := ""
	if report.calendarInfo != "" {
		formattedStr += report.calendarInfo + "\n"
	}
	for k, v := range report.sections {
		//log.Print(k)
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
	report.calendarInfo = getCalendarInfo(day)
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
		var parsedResponse interface{}
		if err = json.Unmarshal(contents, &parsedResponse); err != nil {
			log.Print("Error", err)
		}
		query := parsedResponse.(map[string]interface{})["query"]
		if len(query.(map[string]interface{})) == 0 {
			return ""
		}
		pages := query.(map[string]interface{})["pages"]
		if l := len(pages.(map[string]interface{})); l == 0 {
			return ""
		} else if l > 1 {
			log.Print("Too many responses - ", l)
			return ""
		}
		// there is only one entry exists
		var content string
		for _, v := range pages.(map[string]interface{}) {
			content = v.(map[string]interface{})["extract"].(string)
		}
		return content
	}
	return ""
}

func GetTodaysReport() string {
	location, _ := time.LoadLocation("Europe/Moscow")
	log.Print(location)
	now := time.Now().In(location)
	fullReport := getWikiReport(&now)
	report, err := parseWikiReport(fullReport)
	report.setCalendarInfo(&now)
	if err != nil {
		log.Print("Error:", err)
		return ""
	}
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

func getCalendarInfo(reportDay *time.Time) string {
	firstLine := getFullDateString(reportDay)
	location, _ := time.LoadLocation("Europe/Moscow")

	now := time.Now().In(location)
	year := time.Date(now.Year(), time.December, 31, 0, 0, 0, 0, location)
	today := now.YearDay()
	full_days := year.YearDay()

	rest := full_days - today
	secondLine := ""
	if rest > 0 {
		secondLine = strconv.Itoa(today) + "-й день года. До конца года " + strconv.Itoa(rest) + " " + GetDayNoun(today)
	} else {
		secondLine = "Завтра уже Новый Год!"
	}

	return firstLine + "\n" + secondLine + "\n"
}



func parseWikiReport(fullReport string) (Report, error) {
	if fullReport == "" {
		return Report{}, errors.New("empty report")
	}
	scanner := bufio.NewScanner(strings.NewReader(fullReport))
	var header1 string
	var header2 string
	re := regexp.MustCompile("В .* церкви")
	report := Report{sections: map[string][]*Section{}}
	var currSection *Section
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "== ") && strings.HasSuffix(line, " ==") {
			header1 = strings.TrimSpace(strings.Trim(line, "=="))
			header2 = ""
		} else if strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " ===") {
			header2 = strings.TrimSpace(strings.Trim(line, "==="))
			currSection = &Section{header: header2, content: []string{}}
			report.sections[header1] = append(report.sections[header1], currSection)
		} else if line == "" {
			continue
		} else {
			if header1 == "" {
				report.common = append(report.common, line)
			} else {
				if header2 != "" {
					if header1 == "Праздники и памятные дни" && header2 == "Религиозные" {
						reMemorial := regexp.MustCompile("память .*")
						line = strings.TrimSpace(line)
						if has := re.MatchString(line); has {
							first := re.FindAllString(line, 2)[0]
							second := re.Split(line, 2)[1]
							currSection.content = append(currSection.content, first)
							if second != "" {
								currSection.content = append(currSection.content, "- "+second)
							}
						} else if has := reMemorial.MatchString(line); has {
							continue
						} else {
							currSection.content = append(currSection.content, "- "+line)
						}
					} else {
						currSection.content = append(currSection.content, "- "+strings.TrimSpace(line))
					}
				}
			}
		}
	}

	for _, v := range report.sections {
		for _, sect := range v {
			if sect.header == "Религиозные" {
				clean := true
				for _, val := range sect.content {
					if has := re.MatchString(val); !has {
						clean = false
					}
				}
				if clean {
					sect.content = []string{}
				}
			}
		}
	}

	return report, nil
}

