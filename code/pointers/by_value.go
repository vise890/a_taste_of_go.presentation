package main

import (
	"fmt"
)

// START OMIT
type Foo struct {
	Bar int
}

// f is passed by value, and Bar is updated in the copy
func (f Foo) UselessSetBar(newBar int) {
	f.Bar = newBar
}

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, f.Bar = %d\n", f.Bar)

	// when we pass by copy:
	f.UselessSetBar(5)

	// ..the change doesn't affect the original
	fmt.Printf("After f.UselessSetBar(5), f.Bar = %d\n", f.Bar)
}

// END OMIT
