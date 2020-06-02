package dog

import "testing"

func TestYears(t *testing.T) {
	got := Years(10)
	expected := 70

	if got != expected {
		t.Error("Years(10) Expected: ", expected, "Got: ", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(10)
	expected := 70

	if got != expected {
		t.Error("YearsTwo(10) Expected: ", expected, "Got: ", got)
	}
}

func ExampleYears() {
	Years(10)
	//output 70
}

func ExampleYearsTwo() {
	YearsTwo(10)
	//output 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
