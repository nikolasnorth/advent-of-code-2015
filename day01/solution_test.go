package day01

import "testing"

func TestPart01(t *testing.T) {
	expect := 74
	result, err := part01("./input")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected: %d but got %d", expect, result)
	}
}
