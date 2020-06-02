/*
https://blog.golang.org/defer-panic-and-recover

defer ---> Close right after opening it! Guranteed closing in any case (except os.exit(1) or fatalPrint)
	 ---> args evaluated when the defer statement is evaluated (so you get expected result! close what you opened even if var changes overtime)
	 ----> LIFO order (each defer causes stack entry!)
	 ---->

	 open()
	 defer close()
	 lock()
	 defer unlock()

Panic --> Like throw (stops normal execution of function, executes defered statements, returns error to calling function and so on...)
Recover---> catch + finally panic, fix error, resume normal execution
*/

package main

import "fmt"

func main() {
	//foo() //BUG: Try this here, foo() panics but no arragement is done in main to recover, yet!

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovering.... in main... ", r)
		}
	}()

	fmt.Println("Calling foo.")
	foo(0)
	fmt.Println("Returned normally from foo().")
}

func foo(i int) {
	if i > 3 {
		fmt.Println("Panicking...")
		panic("Panicking from foo() as it reaches to: 4")
	}

	defer fmt.Println("Defer in foo: ", i)
	fmt.Println("foo: ", i)
	foo(i + 1)
}
