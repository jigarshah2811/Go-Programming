package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main()  {
	goRoutineExample()
	channelExample()
}

func channelExample()  {
	channel := make(chan int, 2)

	go func() {
		for i:=0; i<10; i++{
			channel <- i
			t := time.Now()
			fmt.Printf("routine 1: produced %d at: %s", i, t)
		}
	}()

	go func() {
		for {
			i := <-channel
			t := time.Now()
			fmt.Printf("routine 2: consumed %d at: %s", i, t)
		}
	}()
}
func goRoutineExample()  {
	wg.Add(1)
	go funcA()
	wg.Wait()
}

func funcA()  {
	/* Best-Practice: to defer at start of the function
	Because if it's end of fun and panic happens in-between, you leave wait!
	 */
	defer wg.Done()
	fmt.Println("In GoRoutine funcA...")
}