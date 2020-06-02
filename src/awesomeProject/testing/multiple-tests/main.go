package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 5 + 6 = ", mySum(4, 5, 6))
}

//mySum A function that takes list of values and returns the sum of all values passed
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
