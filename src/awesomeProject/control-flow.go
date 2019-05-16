package main

import (
	"fmt"
	"strings"
)

// STATIC programming lang, strictly typed
// package scope
var globalY = 100	// DECLARE a variable outside of main (global), use var
var globalZ string  // DECLARE a variable with type string, ASSIGNS ZERO VALUE automatically

/* User-defined type */
type myInt int
var myIntJ myInt

/* iota = auto increment */
const (
	a = iota
	b = iota
	c = iota
)

func main() {
	/* Variables */
	x := 40 // DECLARE a variable & ASSIGN a value
	y := "Jigar Shah" + "10" // a statement made of expression
	fmt.Println("value of x: {0} and y: {1}", x, y)
	fmt.Println("value of globalY: ", globalY)
	globalZ = "Shaken not stirred"
	fmt.Printf("type of globalY: %T and globalZ: %T\n", globalY, globalZ)

	myIntJ = 10
	fmt.Println("User-defined type myIntJ = ", myIntJ)
	myIntJ = myInt(globalY)
	fmt.Println("type conversion myIntJ = ", myIntJ)
	fmt.Printf("Decimal: %d\t Binary: %b\t Hex: %#x\n", myIntJ, myIntJ, myIntJ)

	// _ tells compiler explicitly that I am ignoring return value
	n, _ := fmt.Println("Hello Class...")
	fmt.Println("returnCode: ", n)

	foo("paramString", 2)

	bar()

}

/* Control flow */
/*
	3 types
	1. Sequence
	2. Loop/Iterative
	3. Conditional
*/
func foo(a ... interface{})  {
	// package.element (element can be func, value, constant etc...)
	fmt.Println("This is function foo... accepting variadic params of any type")

	// for loop: init; condition; post {}
	for i := 0; i<10; i++  {
		if i%2 == 0 {
			fmt.Println("This is even number: ", i)
		} else {
			continue
		}
	}

	// for loop on variadic params
	for i:=0; i< len(a); i++  {
		fmt.Println(a[i])
	}

	// for loop range: use to loop over a collection (array/slice/map/channel)
	for i := range a  {
		fmt.Println(a[i])
	}


}

func bar()  {
	fmt.Println("This is function bar...")
	globalY ++
	fmt.Println("value of globalY: ", globalY)
	fmt.Println("value of globalZ: ", globalZ)

	// const iota print
	fmt.Printf("iota example. a: %d, b: %d, c=%d\n", a,b,c)

	// Operators
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)

	// condition
	//f, err := os.Open("control-flow.go")
	//if err != nil {
	//	fmt.Println("file open error: ", err)
	//}

	/* Switch */
	str := "Jigar"
	switch str {
		case "Krupa":
			fmt.Println("This is Krupa")
		case "Jigar", "Parshvi":	// multiple match case
			fmt.Println("This is Jigar or Parshvi")
			if strings.Compare(str, "Jigar") == 0 {
				fmt.Println("Some work for Jigar")
			} else{
				break // break early from switch since no work for Parshvi!
			}
			fmt.Println("after break")
		default:
			fmt.Println("This is default case")
	}

}
