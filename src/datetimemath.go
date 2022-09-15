package ralphred

import (
	"fmt"
	"log"
	"math"
	"strconv"
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
	time.RFC822Z,
	time.RFC850,
	time.UnixDate,
	time.Kitchen,
}

func parseDateTime(token string) (time.Time, bool) {
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

	token := ""
	var resulting_time, no_time = time.Time{}, true
	for _, arg := range args {
		if token == "" {
			token = arg
		} else {
			token = token + " " + arg
		}
		log.Printf("token: %s\n", token)
		if no_time {
			init_time, err := parseDateTime(token)
			if !err {
				resulting_time = init_time
				no_time = false
			}
		}
	}
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
