package ralphred

import "testing"

func assertResponse(t *testing.T, input []string, expected string) {
	t.Helper()
	items, err := convertCommand(input)
	if len(items) != 1 {
		t.Fatal("Didn't get any items back")
	}
	result := items[0].Title
	if err != nil {
		t.Fatalf("Got an error: %s", err)
	}
	if result != expected {
		t.Fatalf("Got %s expected %s", result, expected)
	}
}

func TestTempConvert(t *testing.T) {
	t.Run("NumberWithUnit", func(t *testing.T) {
		assertResponse(t, []string{"2c", "f"}, "35.6f")
	})
	t.Run("CToF", func(t *testing.T) {
		assertResponse(t, []string{"2", "c", "f"}, "35.6f")
	})
	t.Run("FToC", func(t *testing.T) {
		assertResponse(t, []string{"50", "f", "c"}, "10c")
	})
}

func TestTimeConvert(t *testing.T) {
	t.Run("YearsToSeconds", func(t *testing.T) {
		assertResponse(t, []string{"10", "yr", "s"}, "315360000s")
	})
	t.Run("SecondsToYears", func(t *testing.T) {
		assertResponse(t, []string{"63072000", "s", "yr"}, "2yr")
	})
	t.Run("MonthToSeconds", func(t *testing.T) {
		assertResponse(t, []string{"2", "mt", "s"}, "5256000s")
	})
	t.Run("SecondsToMonths", func(t *testing.T) {
		assertResponse(t, []string{"2628000", "s", "mt"}, "1mt")
	})
	t.Run("DaysToSeconds", func(t *testing.T) {
		assertResponse(t, []string{"3", "dy", "s"}, "259200s")
	})
	t.Run("SecondsToDays", func(t *testing.T) {
		assertResponse(t, []string{"518400", "s", "dy"}, "6dy")
	})
	t.Run("HoursToSeconds", func(t *testing.T) {
		assertResponse(t, []string{"7", "hr", "s"}, "25200s")
	})
	t.Run("SecondsToHours", func(t *testing.T) {
		assertResponse(t, []string{"12600", "s", "hr"}, "3.5hr")
	})
	t.Run("MinutesToSeconds", func(t *testing.T) {
		assertResponse(t, []string{"5", "mn", "s"}, "300s")
	})
	t.Run("SecondsToHours", func(t *testing.T) {
		assertResponse(t, []string{"540", "s", "mn"}, "9mn")
	})
	t.Run("SecondsToMilli", func(t *testing.T) {
		assertResponse(t, []string{"2.5", "s", "ms"}, "2500ms")
	})
	t.Run("SecondsToMicro", func(t *testing.T) {
		assertResponse(t, []string{"2.5", "s", "mcs"}, "2500000mcs")
	})
	t.Run("SecondsToNano", func(t *testing.T) {
		assertResponse(t, []string{"2.5", "s", "ns"}, "2500000000ns")
	})
}

func TestDistanceConvert(t *testing.T) {
	t.Run("YardToFeet", func(t *testing.T) {
		assertResponse(t, []string{"2.5", "yd", "ft"}, "7.5ft")
	})
	t.Run("FeetToYards", func(t *testing.T) {
		assertResponse(t, []string{"12", "ft", "yd"}, "4yd")
	})
	t.Run("MileToFeet", func(t *testing.T) {
		assertResponse(t, []string{"1.2", "mi", "ft"}, "6336ft")
	})
	t.Run("FeetToMile", func(t *testing.T) {
		assertResponse(t, []string{"8976", "ft", "mi"}, "1.7mi")
	})
	t.Run("MeterToFeet", func(t *testing.T) {
		assertResponse(t, []string{"20", "m", "ft"}, "65.6ft")
	})
	t.Run("FeetToMeter", func(t *testing.T) {
		assertResponse(t, []string{"5", "ft", "m"}, "1.5m")
	})
}

func TestDigitalConvert(t *testing.T) {
	t.Run("BytesToBits", func(t *testing.T) {
		assertResponse(t, []string{"2.5", "B", "b"}, "20b")
	})
	t.Run("BitsToBytes", func(t *testing.T) {
		assertResponse(t, []string{"12", "b", "B"}, "1.5B")
	})
}

func TestSiPrefixes(t *testing.T) {
	t.Run("Yotta", func(t *testing.T) {
		assertResponse(t, []string{"0.00000025", "Ym", "m"}, "250000000000000000m")
	})
	t.Run("Zetta", func(t *testing.T) {
		assertResponse(t, []string{"0.0000006", "Zm", "m"}, "600000000000000m")
	})
	t.Run("Exa", func(t *testing.T) {
		assertResponse(t, []string{"0.00000042", "Em", "m"}, "420000000000m")
	})
	t.Run("Peta", func(t *testing.T) {
		assertResponse(t, []string{"0.000085", "Pm", "m"}, "85000000000m")
	})
	t.Run("Tera", func(t *testing.T) {
		assertResponse(t, []string{"0.000073", "Tm", "m"}, "73000000m")
	})
	t.Run("Giga", func(t *testing.T) {
		assertResponse(t, []string{"0.000058", "Gm", "m"}, "58000m")
	})
	t.Run("Mega", func(t *testing.T) {
		assertResponse(t, []string{"0.0072", "Mm", "m"}, "7200m")
	})
	t.Run("Kilo", func(t *testing.T) {
		assertResponse(t, []string{"0.0017", "km", "m"}, "1.7m")
	})
	t.Run("Hecto", func(t *testing.T) {
		assertResponse(t, []string{"5.2", "hm", "m"}, "520m")
	})
	t.Run("Deca", func(t *testing.T) {
		assertResponse(t, []string{"6.3", "dam", "m"}, "63.0m")
	})
	t.Run("Deci", func(t *testing.T) {
		assertResponse(t, []string{"36", "dm", "m"}, "3.6m")
	})
	t.Run("Centi", func(t *testing.T) {
		assertResponse(t, []string{"92", "cm", "m"}, "0.9m")
	})
	t.Run("Milli", func(t *testing.T) {
		assertResponse(t, []string{"4300", "mm", "m"}, "4.3m")
	})
	t.Run("Micro", func(t *testing.T) {
		assertResponse(t, []string{"870000", "mcm", "m"}, "0.9m")
	})
	t.Run("Nano", func(t *testing.T) {
		assertResponse(t, []string{"540000000", "nm", "m"}, "0.5m")
	})
	t.Run("Pico", func(t *testing.T) {
		assertResponse(t, []string{"6900000000000", "pm", "m"}, "6.9m")
	})
	t.Run("Femto", func(t *testing.T) {
		assertResponse(t, []string{"850000000000000", "fm", "m"}, "0.9m")
	})
	t.Run("Atto", func(t *testing.T) {
		assertResponse(t, []string{"3300000000000000000", "am", "m"}, "3.3m")
	})
	t.Run("Zepto", func(t *testing.T) {
		assertResponse(t, []string{"720000000000000000000", "zm", "m"}, "0.7m")
	})
	t.Run("Yocto", func(t *testing.T) {
		assertResponse(t, []string{"610000000000000000000000", "ym", "m"}, "0.6m")
	})
}

func TestDigitalPrefixes(t *testing.T) {
	t.Run("Yotta", func(t *testing.T) {
		assertResponse(t, []string{"0.00000025", "Yb", "b"}, "250000000000000000b")
	})
	t.Run("Zetta", func(t *testing.T) {
		assertResponse(t, []string{"0.0000006", "Zb", "b"}, "600000000000000b")
	})
	t.Run("Exa", func(t *testing.T) {
		assertResponse(t, []string{"0.00000042", "Eb", "b"}, "420000000000b")
	})
	t.Run("Peta", func(t *testing.T) {
		assertResponse(t, []string{"0.000085", "Pb", "b"}, "85000000000b")
	})
	t.Run("Tera", func(t *testing.T) {
		assertResponse(t, []string{"0.000073", "Tb", "b"}, "73000000b")
	})
	t.Run("Giga", func(t *testing.T) {
		assertResponse(t, []string{"0.000058", "Gb", "b"}, "58000b")
	})
	t.Run("Mega", func(t *testing.T) {
		assertResponse(t, []string{"0.0072", "Mb", "b"}, "7200b")
	})
	t.Run("Kilo", func(t *testing.T) {
		assertResponse(t, []string{"0.0017", "kb", "b"}, "1.7b")
	})
	t.Run("Hecto", func(t *testing.T) {
		assertResponse(t, []string{"5.2", "hb", "b"}, "520b")
	})
	t.Run("Deca", func(t *testing.T) {
		assertResponse(t, []string{"6.3", "dab", "b"}, "63b")
	})
	t.Run("Yobi", func(t *testing.T) {
		assertResponse(t, []string{".00000000003", "Yib", "b"}, "36267774588438.9b")
	})
	t.Run("Zebi", func(t *testing.T) {
		assertResponse(t, []string{".000000005", "Zib", "b"}, "5902958103587.1b")
	})
	t.Run("Exbi", func(t *testing.T) {
		assertResponse(t, []string{".0000008", "Eib", "b"}, "922337203685.5b")
	})
	t.Run("Pebi", func(t *testing.T) {
		assertResponse(t, []string{".0000023", "Pib", "b"}, "2589569785.7b")
	})
	t.Run("Tebi", func(t *testing.T) {
		assertResponse(t, []string{".00045", "Tib", "b"}, "494780232.5b")
	})
	t.Run("Gibi", func(t *testing.T) {
		assertResponse(t, []string{".0071", "Gib", "b"}, "7623567.0b")
	})
	t.Run("Mebi", func(t *testing.T) {
		assertResponse(t, []string{".054", "Mib", "b"}, "56623.1b")
	})
	t.Run("Kibi", func(t *testing.T) {
		assertResponse(t, []string{"6.7", "Kib", "b"}, "6860.8b")
	})
}
