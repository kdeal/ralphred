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
