package ralphred

import "testing"

func assertResponse(t *testing.T, input []string, expected string) {
	t.Helper()
	items, err := convertCommand(input)
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
