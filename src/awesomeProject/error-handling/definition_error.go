package main

import (
	"errors"
	"log"
)

/*ErrNoSqrtForNegatives ERROR HANDLING
Built in: https://godoc.org/builtin#error	 Just another Type, error is built-in as int/string
https://blog.golang.org/error-handling-and-go

Go does not support try->catch->finally idiom instead
1. Check for error right after code as common error not exceptional
2. Create new error type	(https://godoc.org/errors) implement struct extending interface Error

*/
var ErrNoSqrtForNegatives = errors.New("Can't get sqrt of a negative number")

func main() {
	output, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%T\n", ErrNoSqrtForNegatives)
	log.Println("sqrt of num is: ", output)
}

func sqrt(num float64) (float64, error) {
	if num < 0 {

		return 0, ErrNoSqrtForNegatives
	}

	return 42, nil
}
