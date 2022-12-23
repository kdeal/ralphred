package ralphred

import (
	"regexp"
	"strings"
)

var numberWithUnitRegex = regexp.MustCompile("^(?P<number>-?[0-9.]+)(?P<remaining>[^0-9.]+)")

func splitUnitFromNumber(args []string) []string {
	new_args := []string{}
	for _, str := range args {
		match := numberWithUnitRegex.FindStringSubmatch(str)
		if match == nil {
			new_args = append(new_args, str)
		} else {
			number := match[numberWithUnitRegex.SubexpIndex("number")]
			remaining := match[numberWithUnitRegex.SubexpIndex("remaining")]
			new_args = append(new_args, number, remaining)
		}
	}
	return new_args
}

func queryMatches(test_string string, terms []string) bool {
	for _, term := range terms {
		if !strings.Contains(test_string, term) {
			return false
		}
	}
	return true
}
