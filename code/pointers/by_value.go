package main

import (
	"fmt"
)

// START OMIT
type Foo struct {
	Bar int
}

// f is passed by value, and Bar is updated in the copy
func UselessSetBar(f Foo, newBar int) {
	f.Bar = newBar
}

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, Bar = %d\n", f.Bar)

	// when we pass by copy:
	UselessSetBar(f, 5)

	// ..the change doesn't affect the original
	fmt.Printf("After UselessSetBar(f, 5), Bar = %d\n", f.Bar)
}

// END OMIT
