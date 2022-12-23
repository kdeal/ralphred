package ralphred

import (
	"reflect"
	"testing"
)

func TestSplitNumbersAndText(t *testing.T) {
	assertResult := func(input []string, expected []string) {
		result := splitUnitFromNumber(input)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Got %s expected %s", result, expected)
		}
	}
	t.Run("EmptyArgs", func(t *testing.T) {
		assertResult([]string{}, []string{})
	})
	t.Run("NoSplit", func(t *testing.T) {
		assertResult([]string{"a", "b"}, []string{"a", "b"})
	})
	t.Run("OneSplit", func(t *testing.T) {
		assertResult([]string{"1a", "b"}, []string{"1", "a", "b"})
	})
	t.Run("FloatSplit", func(t *testing.T) {
		assertResult([]string{"1.2a"}, []string{"1.2", "a"})
	})
	t.Run("NumberAfterText", func(t *testing.T) {
		assertResult([]string{"1a", "b2"}, []string{"1", "a", "b2"})
	})
	t.Run("NegativeNumber", func(t *testing.T) {
		assertResult([]string{"-1a"}, []string{"-1", "a"})
	})
	t.Run("FloatNumberNoUnit", func(t *testing.T) {
		assertResult([]string{"1.2", "a"}, []string{"1.2", "a"})
	})
}

func TestQueryMatches(t *testing.T) {
	assertMatch := func(str string, terms []string, matches bool) {
		result := queryMatches(str, terms)
		if result != matches {
			if matches {
				t.Fatal("Query didn't match when it should")
			} else {
				t.Fatalf("Query matched when it shouldn't")
			}
		}
	}
	t.Run("EmptyString", func(t *testing.T) {
		assertMatch("", []string{"test"}, false)
	})
	t.Run("NoSearchString", func(t *testing.T) {
		assertMatch("test", []string{}, true)
	})
	t.Run("OneSearchStringMatches", func(t *testing.T) {
		assertMatch("test", []string{"te"}, true)
	})
	t.Run("OneSearchStringNoMatch", func(t *testing.T) {
		assertMatch("test", []string{"oue"}, false)
	})
	t.Run("MultiSearchStringNoMatch", func(t *testing.T) {
		assertMatch("test", []string{"te", "st"}, true)
	})
	t.Run("MultiSearchStringOneMatch", func(t *testing.T) {
		assertMatch("test", []string{"te", "oue"}, false)
	})
}
