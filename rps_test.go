package rps

import (
	"reflect"
	"testing"
)

func reset() {
	Set(defaultChoices)
}

func ContainsString(arr []string, element string) (isContained bool) {
	reset()

	for _, v := range arr {
		if reflect.DeepEqual(v, element) {
			isContained = true
			return
		}
	}
	return
}

func TestRps(t *testing.T) {
	choice := Roll(choices)

	if !ContainsString(choices, choice) {
		t.Errorf("Expected %s to contain %s", choices, choice)
	}

	choice = Roll()

	if !ContainsString(choices, choice) {
		t.Errorf("Expected %s to contain %s", choices, choice)
	}
}

func TestCustomChoices(t *testing.T) {
	custom := []string{
		"foo",
		"bar",
		"baz",
	}

	choice := Roll(custom)

	if !ContainsString(custom, choice) {
		t.Errorf("Expected %s to contain %s", custom, choice)
	}
}

func TestSet(t *testing.T) {
	custom := []string{
		"foo",
		"bar",
		"baz",
	}
	Set(custom)

	choice := Roll()
	if !ContainsString(custom, choice) {
		t.Errorf("Expected %s to contain %s", custom, choice)
	}

	reset()
}
