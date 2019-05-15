package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(input)

	fmt.Printf("Factorial of %d is %d", num, factorial(num))

}

func factorial(num int) int  {
	var fact [100]int

	fact[0] = 0
	fact[1] = 1

	for i := 2; i <= num; i++ {
		fact[i] = fact[i-1] * fact[i-2]
	}

	return fact[num]
}