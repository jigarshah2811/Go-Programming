package main

import "fmt"

func main()  {
	class := `https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#`
	fmt.Println("Go exercise from class: ", class)

	Ninja3()
}

func Ninja3()  {
	fmt.Println("Ex1: loop")
	for i := 1; i<=10; i++ {
		fmt.Printf("%d,", i)
	}

	fmt.Println("Ex2: loop within loop")
	for c := 'A'; c <= 'Z'; c++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", c)
		}
	}

	fmt.Println("Ex6: Ex7: loop with if else")
	for i:=0; i<10; i++ {
		if i%2 == 0 {
			fmt.Printf("Even: %v\t", i)
		} else {
			fmt.Printf("Odd: %v\t", i)
		}
	}

	fmt.Println("Ex8: Switch statement")
	var exp string
	exp = "+"
	switch exp {
		case "+", "-", "*", "/", "%":
			fmt.Printf("%s is Operator", exp)
		default:
			fmt.Println("%s is Operand", exp)
	}

	fmt.Println("Ex10: Logical operators, booleans")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
