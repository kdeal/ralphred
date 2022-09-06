package ralphred

import (
	"fmt"
	"strconv"
)

const (
	Temperature = "temperature"
	Distance    = "distance"
)

type Unit struct {
	Name     string
	Symbol   string
	Type     string
	ToBase   func(float64) float64
	FromBase func(float64) float64
}

var units []Unit = []Unit{
	// Temperature Units - base is celsius
	{
		Name:   "celsius",
		Symbol: "c",
		Type:   Temperature,
		ToBase: func(current float64) float64 {
			return current
		},
		FromBase: func(base float64) float64 {
			return base
		},
	},
	{
		Name:   "fahrenheit",
		Symbol: "f",
		Type:   Temperature,
		ToBase: func(current float64) float64 {
			return (current - 32) * 5 / 9
		},
		FromBase: func(base float64) float64 {
			return (base * 9 / 5) + 32
		},
	},
	{
		Name:   "Kelvin",
		Symbol: "k",
		Type:   Temperature,
		ToBase: func(current float64) float64 {
			return current - 273.15
		},
		FromBase: func(base float64) float64 {
			return base + 273.15
		},
	},
	// Distance Units - base is feet
	{
		Name:   "feet",
		Symbol: "ft",
		Type:   Distance,
		ToBase: func(current float64) float64 {
			return current
		},
		FromBase: func(base float64) float64 {
			return base
		},
	},
	{
		Name:   "yard",
		Symbol: "yd",
		Type:   Distance,
		ToBase: func(current float64) float64 {
			return current * 3
		},
		FromBase: func(base float64) float64 {
			return base / 3
		},
	},
	{
		Name:   "mile",
		Symbol: "mi",
		Type:   Distance,
		ToBase: func(current float64) float64 {
			return current * 5280
		},
		FromBase: func(base float64) float64 {
			return base / 5280
		},
	},
	{
		Name:   "meter",
		Symbol: "m",
		Type:   Distance,
		ToBase: func(current float64) float64 {
			return current * 3.280839895
		},
		FromBase: func(base float64) float64 {
			return base / 3.280839895
		},
	},
}

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

	var to_unit Unit
	var from_unit Unit

	for _, unit := range units {
		if to_unit_str == unit.Symbol {
			to_unit = unit
		}
		if from_unit_str == unit.Symbol {
			from_unit = unit
		}
	}

	result := 0.0
	if from_unit.Name == "" && to_unit.Name == "" {
		errMsg := fmt.Sprintf("The units supplied aren't supported \"%s\" and \"%s\"", to_unit_str, from_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Name == "" {
		errMsg := fmt.Sprintf("The unit \"%s\" isn't supported", from_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if to_unit.Name == "" {
		errMsg := fmt.Sprintf("The unit \"%s\" isn't supported", to_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Type != to_unit.Type {
		errMsg := fmt.Sprintf("Unable to convert \"%s\" to \"%s\"", from_unit_str, to_unit_str)
		errorAlfredResponse(errMsg).Print()
		return
	} else if from_unit.Name == to_unit.Name {
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
				Title:        fmt.Sprintf("%.1f %s", result, to_unit.Symbol),
				Subtitle:     "",
				Arg:          []string{resultStr},
				Autocomplete: resultStr,
			},
		},
	}
	resp.Print()
}
