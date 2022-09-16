package ralphred

import (
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
			return init_time, false, args[n:]
		}
	}
	return time.Time{}, true, []string{}
}

func dateTimeMathCommand(args []string) {
	if len(args) == 0 {
		resp := AlfredResponse{
			Items: []AlfredItem{
				alfredItemFromString("Input a time", false),
			},
		}
		resp.Print()
		return
	}

	resulting_time, no_time, remaining_args := parseDateTime(args)
	log.Printf("Args left after parsing time: [%s]\n", strings.Join(remaining_args, ", "))

	items := []AlfredItem{}
	if no_time {
		items = []AlfredItem{
			alfredItemFromString("Unable to parse a time", false),
		}
	} else {
		// +1 is for unix timestamp
		items = make([]AlfredItem, len(output_time_formats)+1)
		for i, format := range output_time_formats {
			items[i] = alfredItemFromString(resulting_time.Format(format), false)
		}
		unix := resulting_time.Unix()
		items[len(items)-1] = alfredItemFromString(fmt.Sprintf("%d", unix), false)
	}
	resp := AlfredResponse{
		Items: items,
	}
	resp.Print()
}
