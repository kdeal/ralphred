package ralphred

import (
	"fmt"
	"strconv"
)

func convertCommand(args []string) {
	if len(args) == 0 {
		errorAlfredResponse("Type measurement with unit to start converting").Print()
		return
	}

	measurement, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		errMsg := fmt.Sprintf("Error converting \"%s\" to a number", args[0])
		errorAlfredResponse(errMsg).Print()
		return
	}

	if len(args) == 1 {
		// TODO: List supported units
		errorAlfredResponse("Please specify a unit for the measurement").Print()
		return
	}

	from_unit_str := args[1]

	if len(args) == 2 {
		// TODO: Give list of units to convert to
		errorAlfredResponse("Please specify a unit to convert to").Print()
		return
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
		errMsg := fmt.Sprintf("The units supplied aren't supported \"%s\" and \"%s\"", to_unit_str, from_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Unit.Name == "" {
		errMsg := fmt.Sprintf("The unit \"%s\" isn't supported", from_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if to_unit.Unit.Name == "" {
		errMsg := fmt.Sprintf("The unit \"%s\" isn't supported", to_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Unit.Type != to_unit.Unit.Type {
		errMsg := fmt.Sprintf("Unable to convert \"%s\" to \"%s\"", from_unit_str, to_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Unit.Name == to_unit.Unit.Name {
		result = measurement
	} else {
		base_value := from_unit.ToBase(measurement)
		result = to_unit.FromBase(base_value)
	}

	resultStr := fmt.Sprintf("%f", result)
	resp := AlfredResponse{
		Items: []AlfredItem{
			{
				UID:          "",
				Title:        fmt.Sprintf("%.1f %s", result, to_unit.Symbol()),
				Subtitle:     "",
				Arg:          []string{resultStr},
				Autocomplete: resultStr,
			},
		},
	}
	resp.Print()
}
