package main

import "testing"

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySum(2, 3, 4, 5, 6)
	}
}

func TestMySum(t *testing.T) {
	expected := 5
	got := mySum(2, 3)
	if got != expected {
		t.Error("Got 2 + 3 = ", got, "Expected: ", expected)
	}

}
