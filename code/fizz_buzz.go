package main

import (
	"fmt"
	"strconv"
)

// FizzBuzz returns:
// - "Fizz" if n is divisible by 3
// - "Buzz" if n is divisible by 5
// - "FizzBuzz" if n is divisible by both 5 and 3
// - n otherwise
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}

}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%2d => %s\n", i, FizzBuzz(i))
	}
}
