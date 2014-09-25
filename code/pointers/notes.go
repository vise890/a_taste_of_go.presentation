package main

import (
	"fmt"
)

// START OMIT
type Foo struct{ Bar int }

func (f *Foo) SetBar(newBar int) { f.Bar = newBar }

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, f.Bar = %d\n", f.Bar)

	// all in one line
	(&f).SetBar(42)
	fmt.Printf("After (&f).SetBar(42), f.Bar = %d\n", f.Bar)
}

// END OMIT
