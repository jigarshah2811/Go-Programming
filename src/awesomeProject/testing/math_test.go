package math

import "testing"

func TestAverage(t *testing.T) {
	actual := Average([]float64{1, 2})
	expected := 1.5

	if actual != expected {
		t.Error("Expected: ", expected, "Actual: ", actual)
	}

}
