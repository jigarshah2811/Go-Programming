package main

import (
	"fmt"
)

func text_le12 (num int) string {
        if (num == 0) {
            return  ""
        }
	lower := make(map[int]string)

	lower[1] = "one"
	lower[2] = "two"
	lower[3] = "three"
	lower[4] = "four"
	lower[5] = "five"
	lower[6] = "six"
	lower[7] = "seven"
	lower[8] = "eight"
	lower[9] = "nine"
	lower[10] = "ten"
	lower[11] = "eleven"
	lower[12] = "twelve"

        return (lower[num])
}

func tenprefix (num int) string {
	ten := make(map[int]string)

        ten[2] = "twent"
        ten[3] = "thirt"
        ten[4] = "fourt"
        ten[5] = "fift"
        ten[6] = "sixt"
        ten[7] = "sevent"
        ten[8] = "eight"
        ten[9] = "ninet"

        return ten[num]
}

func text_13_19 (num int) string {
        return (tenprefix(num % 10 ) + "een")
}

func text_20_99 (num int) string {

    return ((tenprefix(num / 10) + "y ") +  text_le12(num % 10))
}

func texthundred (num int) string {
    if num <= 12 {
        return text_le12(num)
    } else if num < 20 {
        return text_13_19(num)
    } else {
        return text_20_99(num)
    }
}

func texthousands (num int) string {
    if num < 100 {
        return texthundred(num)
    } else {
        return text_le12(num / 100) + " hundred " + texthundred(num % 100)
    }
}

func main() {
        var i int
        var rem int
        prefix := make(map[int]string)
        prefix[0] = ""
        prefix[1] = "thousand "
        prefix[2] = "million "
        prefix[3] = "billion "
        var stack []string

        _, err := fmt.Scanf("%d", &i)
        if err != nil {
            return
        }
        if i == 0 {
            fmt.Println("Zero")
            return
        }
        count := 0
        for i > 0 {
            rem = i % 1000
            i = i / 1000
            if rem > 0 {
                stack = append(stack, texthousands(rem) + " "  + prefix[count])
            }
            count++
        }
        for len(stack) > 0 {
            n := len(stack) - 1 
            fmt.Print(stack[n])
            stack = stack[:n]
        }
}

