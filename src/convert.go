package ralphred

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func convertCommand(args []string) ([]AlfredItem, error) {
	if len(args) == 0 {
		return []AlfredItem{}, errors.New("Type measurement with unit to start converting")
	}

	args = splitUnitFromNumber(args)

	measurement, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		return []AlfredItem{}, fmt.Errorf("Error converting \"%s\" to a number", args[0])
	}

	if len(args) == 1 {
		// TODO: List supported units
		return []AlfredItem{}, errors.New("Please specify a unit for the measurement")
	}

	from_unit_str := args[1]

	if len(args) == 2 {
		// TODO: Give list of units to convert to
		return []AlfredItem{}, errors.New("Please specify a unit to convert to")
	}

	to_unit_str := args[2]

	var to_unit MatchedUnit
	var from_unit MatchedUnit

	for _, unit := range units {
		to_matched_unit, to_matched := unit.matchesString(to_unit_str)
		if to_matched {
			to_unit = to_matched_unit
		}
		from_matched_unit, from_matched := unit.matchesString(from_unit_str)
		if from_matched {
			from_unit = from_matched_unit
		}
	}

	result := 0.0
	if from_unit.Unit.Name == "" && to_unit.Unit.Name == "" {
		return []AlfredItem{}, fmt.Errorf("The units supplied aren't supported \"%s\" and \"%s\"", to_unit_str, from_unit_str)
	} else if from_unit.Unit.Name == "" {
		return []AlfredItem{}, fmt.Errorf("The unit \"%s\" isn't supported", from_unit_str)
	} else if to_unit.Unit.Name == "" {
		return []AlfredItem{}, fmt.Errorf("The unit \"%s\" isn't supported", to_unit_str)
	} else if from_unit.Unit.Type != to_unit.Unit.Type {
		return []AlfredItem{}, fmt.Errorf("Unable to convert \"%s\" to \"%s\"", from_unit_str, to_unit_str)
	} else if from_unit.Symbol() == to_unit.Symbol() {
		result = measurement
	} else {
		base_value := from_unit.ToBase(measurement)
		result = to_unit.FromBase(base_value)
	}

	displayStr := fmt.Sprintf("%.1f%s", result, to_unit.Symbol())
	resultStr := fmt.Sprintf("%f", result)

	if result == math.Trunc(result) {
		displayStr = fmt.Sprintf("%d%s", int64(result), to_unit.Symbol())
		resultStr = fmt.Sprintf("%d", int64(result))
	}

	resp := []AlfredItem{
		{
			UID:          "",
			Title:        displayStr,
			Subtitle:     "",
			Arg:          []string{resultStr},
			Autocomplete: resultStr,
		},
	}
	return resp, nil
}
