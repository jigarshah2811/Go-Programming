package main

import (
	"fmt"
)

func main()  {
	sliceExample()
}

func sliceExample()  {
	// COMPOSIT LITERAL .... type{values}
	mySlice := []int{1, 2, 3, 4, 5}

	// iterate through slice
	for i, v := range mySlice {
		fmt.Printf("Index: %d \t Value: %d\n", i,v)
	}

	// slicing a slice
	fmt.Printf("Slicing from position %d to %d\n", 1, 3)
	fmt.Println(mySlice[1:3])

	// append to a slice
	fmt.Printf("appending to a slice...")
	x := []int{6, 7, 8}
	mySlice = append(mySlice, x...) // ... AFTER means replace all values right here!
	fmt.Println(mySlice)

	// deleting from a slice
	fmt.Printf("deleting from a slice...")
	mySlice = append(mySlice[:2], mySlice[4:]...) // all till 2 and all after 4
	fmt.Println(mySlice)

	// create slice using built-in make (This is FIXED ARRAY)
	fmt.Println("creating slice using built-in make")
	mySlice = make([]int, 10, 11)
	fmt.Println(mySlice)
	fmt.Println("assigning element in slice")
	mySlice[0] = 1
	mySlice[9] = 9
	fmt.Println(mySlice)
	fmt.Printf("slice len=%d, capacity=%d", len(mySlice), cap(mySlice))
	// mySlice[10] = 10 // This won't work since len is limited to 10
	fmt.Println("appending element in slice")
	mySlice = append(mySlice, 11)
	fmt.Println(mySlice)
	fmt.Printf("slice len=%d, capacity=%d", len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 12) // Now is ACTUAL RESIZE happening, capacity doubled!
	fmt.Println(mySlice)
	fmt.Printf("slice len=%d, capacity=%d", len(mySlice), cap(mySlice))



}