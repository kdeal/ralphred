package ralphred

import (
	"testing"
	"time"
)

const DateLayout string = "2006-01-02"

func assertTime(t *testing.T, args []string, expected_time time.Time) {
	t.Helper()
	items, err := dateTimeMathCommand(args)
	if err != nil {
		t.Fatalf("Error %s when getting next time", err.Error())
	}
	var new_time time.Time
	for _, item := range items {
		if item.UID == "RFC3339Milli" {
			new_time, _ = time.Parse("2006-01-02T15:04:05.000Z07:00", item.Title)
			break
		}
	}
	if new_time != expected_time {
		t.Fatalf("Calculated wrong time. Got %s expected %s", new_time, expected_time)
	}
}

func TestFindWeekday(t *testing.T) {
	assertWeekdayTime := func(t *testing.T, weekday string, operation WeekdayOperation, expected_time time.Time) {
		t.Helper()
		args := []string{"2022-09-21", string(operation), weekday}
		assertTime(t, args, expected_time)
	}

	t.Run("NextDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-24")
		assertWeekdayTime(t, "Saturday", NextWeekday, expected_time)
	})
	t.Run("NextDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-26")
		assertWeekdayTime(t, "Monday", NextWeekday, expected_time)
	})
	t.Run("NextSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-28")
		assertWeekdayTime(t, "wednesday", NextWeekday, expected_time)
	})
	t.Run("PrevDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-17")
		assertWeekdayTime(t, "Saturday", PrevWeekday, expected_time)
	})
	t.Run("PrevDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-19")
		assertWeekdayTime(t, "Monday", PrevWeekday, expected_time)
	})
	t.Run("PrevSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-14")
		assertWeekdayTime(t, "wednesday", PrevWeekday, expected_time)
	})
	t.Run("ThisDiffDayLater", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-24")
		assertWeekdayTime(t, "Saturday", ThisWeekday, expected_time)
	})
	t.Run("ThisDiffDayEarlier", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-19")
		assertWeekdayTime(t, "Monday", ThisWeekday, expected_time)
	})
	t.Run("ThisSameDay", func(t *testing.T) {
		expected_time, _ := time.Parse(DateLayout, "2022-09-21")
		assertWeekdayTime(t, "wednesday", ThisWeekday, expected_time)
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

func TestAdd(t *testing.T) {
	t.Run("AddASecond", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T00:00:01Z")
		assertTime(t, []string{"2022-09-21", "+", "1second"}, test_time)
	})
	t.Run("AddFractionalSeconds", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T00:00:01.5Z")
		assertTime(t, []string{"2022-09-21", "+", "1.5seconds"}, test_time)
	})
	t.Run("AddAMinute", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T00:01:00Z")
		assertTime(t, []string{"2022-09-21", "+", "1minute"}, test_time)
	})
	t.Run("AddFractionalMinutes", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T00:02:30Z")
		assertTime(t, []string{"2022-09-21", "+", "2.5minutes"}, test_time)
	})
	t.Run("AddAHour", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T01:00:00Z")
		assertTime(t, []string{"2022-09-21", "+", "1hour"}, test_time)
	})
	t.Run("AddFractionalHours", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-21T01:30:00Z")
		assertTime(t, []string{"2022-09-21", "+", "1.5hours"}, test_time)
	})
	t.Run("AddADay", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-22")
		assertTime(t, []string{"2022-09-21", "+", "1day"}, test_time)
	})
	t.Run("AddFractionalDays", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-22T12:00:00Z")
		assertTime(t, []string{"2022-09-21", "+", "1.5days"}, test_time)
	})
	t.Run("AddAWeek", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-28")
		assertTime(t, []string{"2022-09-21", "+", "1week"}, test_time)
	})
	t.Run("AddFractionalWeek", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-09-24T12:00:00Z")
		assertTime(t, []string{"2022-09-21", "+", ".5weeks"}, test_time)
	})
	t.Run("AddAMonth", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-10-21")
		assertTime(t, []string{"2022-09-21", "+", "1month"}, test_time)
	})
	t.Run("AddAYear", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2023-09-21")
		assertTime(t, []string{"2022-09-21", "+", "1year"}, test_time)
	})
	t.Run("AddFractionalYear", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2024-03-21")
		assertTime(t, []string{"2022-09-21", "+", "1.5year"}, test_time)
	})
	t.Run("AddBadFractionalYear", func(t *testing.T) {
		res, err := dateTimeMathCommand([]string{"2022-03-21", "+", "1.3year"})
		if err == nil {
			t.Fatalf("Expected an error, but got: %s", res)
		}
	})
	t.Run("AddMultipleOfAUnit", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-23")
		assertTime(t, []string{"2022-09-21", "+", "2day"}, test_time)
	})
	t.Run("AddMultipleUnits", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-30")
		assertTime(t, []string{"2022-09-21", "+", "2day", "1week"}, test_time)
	})
	t.Run("NotANumber", func(t *testing.T) {
		res, err := dateTimeMathCommand([]string{"2022-09-21", "+", "one", "year"})
		if err == nil {
			t.Fatalf("Expected an error, but got: %s", res)
		}
	})
	t.Run("NotEnoughArgs", func(t *testing.T) {
		res, err := dateTimeMathCommand([]string{"2022-09-21", "+", "year"})
		if err == nil {
			t.Fatalf("Expected an error, but got: %s", res)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("SubADay", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-20")
		assertTime(t, []string{"2022-09-21", "-", "1day"}, test_time)
	})
	t.Run("SubANegativeDay", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2022-09-22")
		assertTime(t, []string{"2022-09-21", "-", "-1day"}, test_time)
	})
	t.Run("SubARounding", func(t *testing.T) {
		test_time, _ := time.Parse(DateLayout, "2021-03-21")
		assertTime(t, []string{"2022-09-21", "-", "1.5year"}, test_time)
	})
}

func TestFloor(t *testing.T) {
	t.Run("FloorToMinute", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-05T05:05:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "minute"}, test_time)
	})
	t.Run("FloorToHour", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-05T05:00:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "hour"}, test_time)
	})
	t.Run("FloorToDay", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-05T00:00:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "day"}, test_time)
	})
	t.Run("FloorToWeek", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-02T00:00:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "week"}, test_time)
	})
	t.Run("FloorToMonth", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-01T00:00:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "month"}, test_time)
	})
	t.Run("FloorToYear", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-01-01T00:00:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "floor", "year"}, test_time)
	})
	t.Run("StartOfMinute", func(t *testing.T) {
		test_time, _ := time.Parse(time.RFC3339, "2022-05-05T05:05:00Z")
		assertTime(t, []string{"2022-05-05T05:05:05Z", "start", "minute"}, test_time)
	})
}
