package main

import (
	"fmt"
	"math/rand"
)

const maxTemp int = 50

// START OMIT
type City struct {
	Name    string
	Country string
}

func (c City) String() string {
	return c.Name + "," + c.Country
}

func (c City) CurrentTemp() int {
	return rand.Intn(maxTemp)
}

func main() {
	bologna := City{
		Name:    "Bologna",
		Country: "Italy",
	}
	fmt.Printf("Temperature in %s: %dC", bologna.String(), bologna.CurrentTemp())
}

// END OMIT
