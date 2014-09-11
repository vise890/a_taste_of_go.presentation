package main

import "fmt"

// START OMIT
func fib(n int) int {
	// n <= 0 not handled..
	if n == 1 || n == 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// go is boring and old school with these things:
	for i := 1; i < 20; i++ {
		fmt.Println(fib(i))
	}
}

// END OMIT
