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
	time.UnixDate,
	time.RFC3339Nano,
	time.Kitchen,
	time.ANSIC,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
}

var output_time_formats []string = []string{
	time.RFC3339,
	time.RFC1123,
	"2006-01-02",
	"15:04:05",
	"Jan _2, 2006",
	time.Kitchen,
}

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
	for i, format := range output_time_formats {
		items[i] = alfredItemFromString(resulting_time.Format(format), false)
	}
	unix := resulting_time.Unix()
	items[len(items)-1] = alfredItemFromString(fmt.Sprintf("%d", unix), false)

	return items, nil
}
