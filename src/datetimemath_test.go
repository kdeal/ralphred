package ralphred

import (
	"testing"
	"time"
)

const DateLayout string = "2006-01-02"

func TestFindWeekday(t *testing.T) {
	assertTime := func(t *testing.T, weekday string, operation WeekdayOperation, expected_time time.Time) {
		t.Helper()
		args := []string{"2022-09-21", string(operation), weekday}
		items, err := dateTimeMathCommand(args)
		if err != nil {
			t.Fatalf("Error %s when getting next time", err.Error())
		}
		var new_time time.Time
		for _, item := range items {
			if item.UID == "RFC3339" {
				new_time, _ = time.Parse(time.RFC3339, item.Title)
				break
			}
		}
		if new_time != expected_time {
			t.Fatalf("Calculated wrong time. %s expected %s", new_time, expected_time)
		}
	}

	t.Run("NextDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-24")
		assertTime(t, "Saturday", NextWeekday, expected_time)
	})
	t.Run("NextDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-26")
		assertTime(t, "Monday", NextWeekday, expected_time)
	})
	t.Run("NextSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-28")
		assertTime(t, "wednesday", NextWeekday, expected_time)
	})
	t.Run("PrevDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-17")
		assertTime(t, "Saturday", PrevWeekday, expected_time)
	})
	t.Run("PrevDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-19")
		assertTime(t, "Monday", PrevWeekday, expected_time)
	})
	t.Run("PrevSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-14")
		assertTime(t, "wednesday", PrevWeekday, expected_time)
	})
	t.Run("ThisDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-24")
		assertTime(t, "Saturday", ThisWeekday, expected_time)
	})
	t.Run("ThisDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-19")
		assertTime(t, "Monday", ThisWeekday, expected_time)
	})
	t.Run("ThisSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-21")
		assertTime(t, "wednesday", ThisWeekday, expected_time)
	})
	t.Run("InvalidWeekday", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-21")
		_, err := findWeekday(test_time, []string{"wat"}, NextWeekday)
		if err == nil {
			t.Fatalf("Expected error, but didn't get one")
		}
	})
	t.Run("TooManyArgs", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-21")
		_, err := findWeekday(test_time, []string{"wat", "wat"}, NextWeekday)
		if err == nil {
			t.Fatalf("Expected error, but didn't get one")
		}
	})
	t.Run("NoArgs", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-21")
		_, err := findWeekday(test_time, []string{}, NextWeekday)
		if err == nil {
			t.Fatalf("Expected error, but didn't get one")
		}
	})
}
