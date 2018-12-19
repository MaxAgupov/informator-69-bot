package wiki

import (
	"bufio"
	"errors"
	"log"
	"regexp"
	"strings"
)

type Parser struct {
	report      *Report
	header      string
	subheader   string
	filledSlice *[]string
	parser      func(line string)
}

func (parser *Parser) reset() {
	parser.header = ""
	parser.subheader = ""
	parser.filledSlice = nil
	parser.parser = nil
}

func (parser *Parser) setHeader(header string, parserFunc func(line string)) {
	parser.header = header
	parser.subheader = ""
	parser.filledSlice = nil
	parser.parser = parserFunc
}

func (parser *Parser) setSubheader(subheader string) {
	parser.subheader = strings.TrimSpace(subheader)
	parser.filledSlice = nil
}

func (parser *Parser) parseHolidays(line string) {
	line = strings.Trim(line, ".;— ")
	if parser.subheader == "" && !strings.HasPrefix(line, "См. также:") {
		parser.report.holidaysInt = append(parser.report.holidaysInt, line)
		return
	} else if parser.filledSlice == nil && parser.subheader != rlgHolidaysSubheader {
		switch parser.subheader {
		case intHolidaysSubheader:
			parser.filledSlice = &parser.report.holidaysInt
		case locHolidaysSubheader:
			parser.filledSlice = &parser.report.holidaysLoc
		case profHolidaysSubheader:
			parser.filledSlice = &parser.report.holidaysProf
		case nameDaysSubheader:
			parser.filledSlice = &parser.report.nameDays
		default:
			parser.subheader = ""
			return
		}
	} else if parser.subheader == rlgHolidaysSubheader {
		re := regexp.MustCompile("В .* церкви")
		reOrth := regexp.MustCompile("Православие")
		reCath := regexp.MustCompile("Католицизм")
		reOth := regexp.MustCompile("Другие конфессии")
		switch {
		case re.MatchString(line):
			first := re.FindAllString(line, 2)[0]
			line = re.Split(line, 2)[1]
			switch first {
			case "В православной церкви":
				parser.filledSlice = &parser.report.holidaysRlg.orthodox
			case "В католичечкой церкви":
				parser.filledSlice = &parser.report.holidaysRlg.catholicism
			default:
				parser.filledSlice = &parser.report.holidaysRlg.others
			}
			if line == "" {
				return
			}
		case reOrth.MatchString(line):
			parser.filledSlice = &parser.report.holidaysRlg.orthodox
			line = reOrth.Split(line, 2)[1]
			if line == "" {
				return
			}
		case reCath.MatchString(line):
			parser.filledSlice = &parser.report.holidaysRlg.catholicism
			line = reCath.Split(line, 2)[1]
			if line == "" {
				return
			}
		case reOth.MatchString(line):
			parser.filledSlice = &parser.report.holidaysRlg.others
			line = reOth.Split(line, 2)[1]
			if line == "" {
				return
			}
		}
		reApostle := regexp.MustCompile("память апостол.*")
		reMemorial := regexp.MustCompile("память .*")

		if has := reMemorial.MatchString(line); has {
			if has = reApostle.MatchString(line); !has {
				return
			}
		}
	}
	if parser.filledSlice == nil {
		log.Print("Error parsing:", line)
		return
	}
	*parser.filledSlice = append(*parser.filledSlice, line)
}

func (parser *Parser) parseOmens(line string) {
	if parser.filledSlice == nil {
		parser.filledSlice = &parser.report.omens
	}
	lines := strings.Split(line, ".")
	for _, l := range lines {
		line = strings.Trim(l, ". ")
		if line == "" {
			return
		}
		*parser.filledSlice = append(*parser.filledSlice, line)
	}
}

func Parse(fullReport string) (Report, error) {
	report := Report{}
	if fullReport == "" {
		return report, errors.New("empty report")
	}
	scanner := bufio.NewScanner(strings.NewReader(fullReport))
	parser := Parser{report: &report}

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "== ") && strings.HasSuffix(line, " =="):
			switch header := strings.TrimSpace(strings.Trim(line, "==")); header {
			case holidaysHeader, "Праздники":
				parser.setHeader(header, parser.parseHolidays)
			case "События", "Родились", "Скончались":
				parser.reset()
			case "Приметы", "Народный календарь", "Народный календарь и приметы", "Народный календарь, приметы", "Народный календарь, приметы и фольклор Руси":
				parser.setHeader(header, parser.parseOmens)
			default:
				parser.reset()
				log.Print("Extra header:", header)
			}
		case strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " ==="):
			parser.setSubheader(strings.Trim(line, "==="))
		case line == "":
			continue
		default:
			if parser.parser == nil {
				continue
			}
			parser.parser(strings.TrimSpace(line))
		}
	}
	return report, nil
}
