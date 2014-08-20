// +build OMIT

package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	// START OMIT
	sayHello := func(i int) {
		pause := time.Duration(rand.Intn(1e3)) * time.Millisecond
		time.Sleep(pause)
		log.Printf("`Hello!` from goroutine #%d\n", i)
	}

	for i := 1; i <= 5; i++ {
		log.Printf("Launching goroutine #%d\n", i)
		go sayHello(i)
	}
	// END OMIT

	// You should use a WaitGroup, this is a dirty hack to avoid introducing too
	// much at once
	programLifetime := time.Duration(2) * time.Second
	time.Sleep(programLifetime)
}
