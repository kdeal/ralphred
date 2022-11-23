package ralphred

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

var input_time_formats []string = []string{
	time.RFC3339,
	time.RFC1123,
	"2006-01-02",
	time.UnixDate,
	time.RFC3339Nano,
	time.Kitchen,
	time.ANSIC,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
}

var output_time_formats map[string]string = map[string]string{
	"RFC3339":        time.RFC3339,
	"RFC1123":        time.RFC1123,
	"Date":           "2006-01-02",
	"KitchenSeconds": "15:04:05",
	"WrittenDate":    "Jan _2, 2006",
	"Kitchen":        time.Kitchen,
	"RFC3339Milli":   "2006-01-02T15:04:05.000Z07:00",
	"RFC3339Nano":    time.RFC3339Nano,
}

var daysOfWeek = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Sun":       time.Sunday,
	"Monday":    time.Monday,
	"Mon":       time.Monday,
	"Tuesday":   time.Tuesday,
	"Tue":       time.Tuesday,
	"Wednesday": time.Wednesday,
	"Wed":       time.Wednesday,
	"Thursday":  time.Thursday,
	"Thur":      time.Thursday,
	"Friday":    time.Friday,
	"Fri":       time.Friday,
	"Saturday":  time.Saturday,
	"Sat":       time.Saturday,
}

type WeekdayOperation string

const (
	NextWeekday WeekdayOperation = "next"
	PrevWeekday WeekdayOperation = "prev"
	ThisWeekday WeekdayOperation = "this"
)

func parseDateTimeFromToken(token string) (time.Time, bool) {
	if token == "now" {
		return time.Now(), false
	} else if token == "utc" {
		return time.Now().UTC(), false
	}
	for _, layout := range input_time_formats {
		time, err := time.Parse(layout, token)
		if err == nil {
			log.Printf("Parsed time with %s -> %s\n", layout, time)
			return time, false
		}
	}
	seconds, err := strconv.ParseFloat(token, 64)
	if err == nil {
		millis := int64(math.Round(seconds * 1000))
		return time.UnixMilli(millis), false
	}
	return time.Time{}, true
}

func parseDateTime(args []string) (time.Time, bool, []string) {
	token := ""
	for n, arg := range args {
		if token == "" {
			token = arg
		} else {
			token = token + " " + arg
		}
		init_time, err := parseDateTimeFromToken(token)
		if !err {
			return init_time, false, args[n+1:]
		}
	}
	return time.Time{}, true, []string{}
}

func findWeekday(init_time time.Time, args []string, operation WeekdayOperation) (time.Time, error) {
	if len(args) > 1 {
		return init_time, fmt.Errorf("%s only accepts 1 argument", operation)
	} else if len(args) == 0 {
		return init_time, fmt.Errorf("%s only accepts 1 argument", operation)
	}

	// TODO: If it prefix matches only one day then accept it
	weekdayStr := strings.Title(args[0])
	weekday, ok := daysOfWeek[weekdayStr]
	if !ok {
		return init_time, fmt.Errorf("Unrecognized weekday %s", weekdayStr)
	}

	diff := int(weekday - init_time.Weekday())
	if diff <= 0 && operation == NextWeekday {
		diff += 7
	} else if diff >= 0 && operation == PrevWeekday {
		diff -= 7
	}
	return init_time.AddDate(0, 0, diff), nil
}

func addDurationToTime(init_time time.Time, durationStr string) time.Time {
	dur, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Panic("Failed to parse formatted duration.")
	}
	return init_time.Add(dur)
}

func prependNewUnit(remaining float64, unit string, negate bool, args []string) []string {
	if remaining == 0 {
		return args
	}
	if negate {
		remaining *= -1
	}
	remStr := strconv.FormatFloat(remaining, 'f', -1, 64)
	return append([]string{remStr, unit}, args...)
}

func addToTime(init_time time.Time, args []string, negate bool) (time.Time, error) {
	args = splitUnitFromNumber(args)
	for {
		if len(args) == 0 {
			return init_time, nil
		} else if len(args) < 2 {
			return init_time, fmt.Errorf("Unused argument when parsing add units: %s", args)
		}
		var valueStr, unit string
		valueStr, unit, args = args[0], args[1], args[2:]

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return init_time, fmt.Errorf("Expected number, but got %s", valueStr)
		}

		if negate {
			value *= -1
		}

		switch unit {
		// For second, minute, hour we don't use parseDuration directly since
		// it only supports the single character unit
		case "s", "second", "seconds":
			init_time = addDurationToTime(init_time, valueStr+"s")

		case "m", "minute", "minutes":
			init_time = addDurationToTime(init_time, valueStr+"m")

		case "h", "hour", "hours":
			init_time = addDurationToTime(init_time, valueStr+"h")

		case "d", "day", "days":
			valueInt, valueRem := math.Modf(value)
			init_time = init_time.AddDate(0, 0, int(valueInt))
			args = prependNewUnit(24*valueRem, "hour", negate, args)

		case "w", "week", "weeks":
			args = prependNewUnit(7*value, "day", negate, args)

		case "mn", "month", "months":
			valueInt, valueRem := math.Modf(value)
			if valueRem != 0.0 {
				return init_time, errors.New("Fractional months not supported")
			}
			init_time = init_time.AddDate(0, int(valueInt), 0)
			args = prependNewUnit(30*valueRem, "day", negate, args)

		case "y", "year", "years":
			valueInt, valueRem := math.Modf(value)
			init_time = init_time.AddDate(int(valueInt), 0, 0)

			months := 12 * valueRem
			_, monthsRem := math.Modf(months)
			if monthsRem != 0.0 {
				return init_time, errors.New("Fractional years only supported if it results in even months")
			}
			args = prependNewUnit(months, "month", negate, args)
		default:
			return init_time, fmt.Errorf("Unsupported unit: %s", unit)
		}
	}
}

func floorTime(init_time time.Time, args []string) (time.Time, error) {
	if len(args) != 1 {
		return init_time, fmt.Errorf("floor/start expects 1 argument got: %s", args)
	}
	unit := args[0]
	var new_time time.Time
	switch unit {
	case "minute":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			init_time.Hour(),
			init_time.Minute(),
			0,
			0,
			init_time.Location(),
		)
	case "hour":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			init_time.Hour(),
			0,
			0,
			0,
			init_time.Location(),
		)
	case "day":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			0,
			0,
			0,
			0,
			init_time.Location(),
		)
	case "week":
		day_floor, _ := floorTime(init_time, []string{"day"})
		new_time, _ = findWeekday(day_floor, []string{"sunday"}, ThisWeekday)
	case "month":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			1,
			0,
			0,
			0,
			0,
			init_time.Location(),
		)
	case "year":
		new_time = time.Date(
			init_time.Year(),
			1,
			1,
			0,
			0,
			0,
			0,
			init_time.Location(),
		)
	default:
		return init_time, fmt.Errorf("Unsupported unit: %s", unit)
	}
	return new_time, nil
}

func daysIn(month time.Month, year int) int {
	// From: https://www.brandur.org/fragments/go-days-in-month
	// day = 0 means it goes back one day and since we set the month
	// to the next one it gives the last day of the month we want
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func ceilTime(init_time time.Time, args []string) (time.Time, error) {
	if len(args) != 1 {
		return init_time, fmt.Errorf("ceil/end expects 1 argument got: %s", args)
	}
	unit := args[0]
	var new_time time.Time
	switch unit {
	case "minute":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			init_time.Hour(),
			init_time.Minute(),
			59,
			999999999,
			init_time.Location(),
		)
	case "hour":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			init_time.Hour(),
			59,
			59,
			999999999,
			init_time.Location(),
		)
	case "day":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			init_time.Day(),
			23,
			59,
			59,
			999999999,
			init_time.Location(),
		)
	case "week":
		day_floor, _ := ceilTime(init_time, []string{"day"})
		new_time, _ = findWeekday(day_floor, []string{"saturday"}, ThisWeekday)
	case "month":
		new_time = time.Date(
			init_time.Year(),
			init_time.Month(),
			daysIn(init_time.Month(), init_time.Year()),
			23,
			59,
			59,
			999999999,
			init_time.Location(),
		)
	case "year":
		new_time = time.Date(
			init_time.Year(),
			12,
			31,
			23,
			59,
			59,
			999999999,
			init_time.Location(),
		)
	default:
		return init_time, fmt.Errorf("Unsupported unit: %s", unit)
	}
	return new_time, nil
}

type TimeOperation struct {
	Commands []string
	Apply    func(time.Time, []string) (time.Time, error)
}

var identityOperation = TimeOperation{
	Commands: []string{},
	Apply: func(init_time time.Time, args []string) (time.Time, error) {
		// This is just used as the initial operation and should be overwritten
		// by the first token and never have any arguments
		if len(args) > 0 {
			panic("Identity time operation called with arguments")
		}
		return init_time, nil
	},
}

var operations = []TimeOperation{
	{
		Commands: []string{"to", "in"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			convert_to := strings.Join(args, " ")
			switch convert_to {
			case "":
				return init_time, errors.New("Timezone required after \"to\"")
			case "utc":
				init_time = init_time.UTC()
			case "local":
				init_time = init_time.Local()
			default:
				// TODO: Do partial timezone matches
				loc, err := time.LoadLocation(convert_to)
				if err != nil {
					return init_time, fmt.Errorf("Unrecognized timezone: %s", convert_to)
				}
				init_time = init_time.In(loc)
			}
			return init_time, nil
		},
	},
	{
		Commands: []string{"next"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			return findWeekday(init_time, args, NextWeekday)
		},
	},
	{
		Commands: []string{"prev"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			return findWeekday(init_time, args, PrevWeekday)
		},
	},
	{
		Commands: []string{"this"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			return findWeekday(init_time, args, ThisWeekday)
		},
	},
	{
		Commands: []string{"+"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			return addToTime(init_time, args, false)
		},
	},
	{
		Commands: []string{"-"},
		Apply: func(init_time time.Time, args []string) (time.Time, error) {
			return addToTime(init_time, args, true)
		},
	},
	{
		Commands: []string{"floor", "start"},
		Apply:    floorTime,
	},
	{
		Commands: []string{"ceil", "end"},
		Apply:    ceilTime,
	},
}

func makeOperationsMap() map[string]TimeOperation {
	operation_map := make(map[string]TimeOperation)
	for _, operation := range operations {
		for _, command := range operation.Commands {
			operation_map[command] = operation
		}
	}
	return operation_map
}

func adjustTime(init_time time.Time, args []string) (time.Time, error) {
	if len(args) == 0 {
		return init_time, nil
	}

	operation_map := makeOperationsMap()

	// Ensure first token is a valid operation
	if _, exists := operation_map[args[0]]; !exists {
		return init_time, errors.New("First word after time definition must be an operation")
	}

	var err error
	var operation TimeOperation = identityOperation
	var operation_args = []string{}

	for {
		if len(args) == 0 {
			init_time, err = operation.Apply(init_time, operation_args)
			break
		}

		var token string
		token, args = args[0], args[1:]

		new_operation, exists := operation_map[token]
		if exists {
			init_time, err = operation.Apply(init_time, operation_args)
			if err != nil {
				break
			}
			operation = new_operation
			operation_args = []string{}
		} else {
			operation_args = append(operation_args, token)
		}
	}

	if err != nil {
		return init_time, err
	}

	return init_time, nil
}

func dateTimeMathCommand(args []string) ([]AlfredItem, error) {
	if len(args) == 0 {
		items := []AlfredItem{
			alfredItemFromString("Input a time", false),
		}
		return items, nil
	}

	resulting_time, no_time, remaining_args := parseDateTime(args)
	if no_time {
		return []AlfredItem{}, errors.New("Unable to parse a time")
	}

	log.Printf("Args left after parsing time: [%s]\n", strings.Join(remaining_args, ", "))

	if len(remaining_args) > 0 {
		var err error
		resulting_time, err = adjustTime(resulting_time, remaining_args)
		if err != nil {
			return []AlfredItem{}, err
		}
	}

	// +1 is for unix timestamp
	items := make([]AlfredItem, len(output_time_formats)+1)
	index := 0
	for name, format := range output_time_formats {
		formatted_time := resulting_time.Format(format)
		items[index] = AlfredItem{
			UID:          name,
			Title:        formatted_time,
			Arg:          []string{formatted_time},
			Autocomplete: formatted_time,
		}
		index += 1
	}

	unix_str := fmt.Sprintf("%d", resulting_time.Unix())
	items[len(items)-1] = AlfredItem{
		UID:          "UnixTimeStamp",
		Title:        unix_str,
		Arg:          []string{unix_str},
		Autocomplete: unix_str,
	}

	return items, nil
}
