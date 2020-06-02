package main

import "testing"

//TestMySum The test that runs multiple tests to ensure MySum function executes as expected
func TestMySum(t *testing.T) {
	/* Multiple tests - input structure */
	// Create inptu structures
	type test struct {
		data     []int
		expected int
	}

	tests := []test{
		test{data: []int{2, 3}, expected: 5},
		test{data: []int{-1, 0, 1}, expected: 0},
		test{data: []int{20, 30, 40}, expected: 100},
	}

	// execution phase
	var got int
	for _, v := range tests {
		got = mySum(v.data...)
		if got != v.expected {
			t.Error("Input: ", v.data, "got: ", got, "Expected: ", v.expected)
		}

	}
}
