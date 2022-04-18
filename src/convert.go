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

	from_unit := args[1]

	if len(args) == 2 {
		// TODO: Give list of units to convert to
		errorAlfredResponse("Please specify a unit to convert to").Print()
		return
	}

	to_unit := args[2]

	result := 0.0
	if from_unit == "f" && to_unit == "c" {
		result = (measurement - 32) * 5 / 9
	} else if to_unit == "c" && from_unit == "f" {
		result = (measurement * 9 / 5) + 32
	} else {
		errMsg := fmt.Sprintf("Unable to convert from \"%s\" to \"%s\"", to_unit, from_unit)
		errorAlfredResponse(errMsg).Print()
		return
	}

	resultStr := fmt.Sprintf("%f", result)
	resp := AlfredResponse {
		Items: []AlfredItem{
			{
				UID: "",
				Title: fmt.Sprintf("%.1f %s", result, to_unit),
				Subtitle: "",
				Arg: []string{resultStr},
				Autocomplete: resultStr,
			},
		},
	}
	resp.Print()
}
