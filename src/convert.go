package ralphred

import (
	"fmt"
	"math"
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
	Prefixes []Prefix
	ToBase   func(float64) float64
	FromBase func(float64) float64
}

func (u Unit) matchesString(str string) (MatchedUnit, bool) {
	if u.Prefixes == nil && str == u.Symbol {
		return MatchedUnit{Unit: u}, true
	} else if u.Prefixes != nil {
		for _, prefix := range prefixes {
			prefixedSymbol := prefix.Symbol + u.Symbol
			if str == prefixedSymbol {
				return MatchedUnit{Unit: u, Prefix: prefix}, true
			}
		}
	}
	return MatchedUnit{}, false
}

type MatchedUnit struct {
	Unit   Unit
	Prefix Prefix
}

func (u MatchedUnit) Symbol() string {
	return u.Prefix.Symbol + u.Unit.Symbol
}

func (u MatchedUnit) Scale() float64 {
	scale := 1.0
	if u.Unit.Prefixes != nil {
		scale = math.Pow(u.Prefix.Base, u.Prefix.Exponent)
	}
	return scale
}

func (u MatchedUnit) ToBase(measurement float64) float64 {
	scale := u.Scale()
	return u.Unit.ToBase(measurement * scale)
}

func (u MatchedUnit) FromBase(measurement float64) float64 {
	scale := u.Scale()
	return u.Unit.FromBase(measurement) / scale
}

type Prefix struct {
	Name     string
	Symbol   string
	Exponent float64
	Base     float64
}

var prefixes []Prefix = []Prefix{
	{
		Name:     "yotta",
		Symbol:   "Y",
		Exponent: 24,
		Base:     10,
	},
	{
		Name:     "zetta",
		Symbol:   "Z",
		Exponent: 21,
		Base:     10,
	},
	{
		Name:     "exa",
		Symbol:   "E",
		Exponent: 18,
		Base:     10,
	},
	{
		Name:     "peta",
		Symbol:   "P",
		Exponent: 15,
		Base:     10,
	},
	{
		Name:     "tera",
		Symbol:   "T",
		Exponent: 12,
		Base:     10,
	},
	{
		Name:     "giga",
		Symbol:   "G",
		Exponent: 9,
		Base:     10,
	},
	{
		Name:     "mega",
		Symbol:   "M",
		Exponent: 6,
		Base:     10,
	},
	{
		Name:     "kilo",
		Symbol:   "k",
		Exponent: 3,
		Base:     10,
	},
	{
		Name:     "hecto",
		Symbol:   "h",
		Exponent: 2,
		Base:     10,
	},
	{
		Name:     "deca",
		Symbol:   "da",
		Exponent: 1,
		Base:     10,
	},
	{
		Name:     "",
		Symbol:   "",
		Exponent: 0,
		Base:     10,
	},
	{
		Name:     "deci",
		Symbol:   "d",
		Exponent: -1,
		Base:     10,
	},
	{
		Name:     "centi",
		Symbol:   "c",
		Exponent: -2,
		Base:     10,
	},
	{
		Name:     "milli",
		Symbol:   "m",
		Exponent: -3,
		Base:     10,
	},
	{
		Name:     "micro",
		Symbol:   "mc",
		Exponent: -6,
		Base:     10,
	},
	{
		Name:     "nano",
		Symbol:   "n",
		Exponent: -9,
		Base:     10,
	},
	{
		Name:     "pico",
		Symbol:   "p",
		Exponent: -12,
		Base:     10,
	},
	{
		Name:     "femto",
		Symbol:   "f",
		Exponent: -15,
		Base:     10,
	},
	{
		Name:     "atto",
		Symbol:   "a",
		Exponent: -18,
		Base:     10,
	},
	{
		Name:     "zepto",
		Symbol:   "z",
		Exponent: -21,
		Base:     10,
	},
	{
		Name:     "yocto",
		Symbol:   "y",
		Exponent: -24,
		Base:     10,
	},
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
		Name:     "meter",
		Symbol:   "m",
		Type:     Distance,
		Prefixes: prefixes,
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
