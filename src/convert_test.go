package ralphred

import "testing"

func TestConvertCommand(t *testing.T) {
	assertResponse := func(input []string, expected string) {
		items, err := convertCommand(input)
		result := items[0].Title
		if err != nil {
			t.Fatalf("Got an error: %s", err)
		}
		if result != expected {
			t.Fatalf("Got %s expected %s", result, expected)
		}
	}
	t.Run("NumberWithUnit", func(t *testing.T) {
		assertResponse([]string{"2c", "f"}, "35.6f")
	})
	t.Run("CToF", func(t *testing.T) {
		assertResponse([]string{"2", "c", "f"}, "35.6f")
	})
	t.Run("FToC", func(t *testing.T) {
		assertResponse([]string{"50", "f", "c"}, "10.0c")
	})
}
