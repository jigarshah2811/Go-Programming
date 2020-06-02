package main

import "testing"

func TestMySum(t *testing.T) {
	var expected, got int

	expected = 5
	got = mySum(2, 3)
	if got != expected {
		t.Error("Expected: ", expected, "Actual: ", got)
	}

	expected = 14
	got = mySum(4, 5, 6)
	if got != expected {
		t.Error("Expected: ", expected, "Actual: ", got)
	}
}
