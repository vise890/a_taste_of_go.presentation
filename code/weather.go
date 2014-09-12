package main

import (
	"fmt"
	"math/rand"
)

const maxTemp int = 40

// START OMIT
type City struct {
	Name    string
	Country string
}

func (c City) String() string {
	return c.Name + "," + c.Country
}

func (c City) CurrentTemp() float32 {
	return rand.Float32() * 50
}

func main() {
	bologna := City{
		Name:    "Bologna",
		Country: "Italy",
	}
	fmt.Printf("Temperature in %s: %dC", bologna.String(), bologna.CurrentTemp())
}

// END OMIT
