package main

import (
	"fmt"
)

// START OMIT
type Foo struct{ Bar int }

// f is passed by reference:
//   - the argument `f` is of type *Foo (read pointer to Foo)
//   - this time bar will be updated in the Foo that we pass in
func SetBar(f *Foo, newBar int) { f.Bar = newBar }

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, Bar = %d\n", f.Bar)

	// SetBar(f, 5) // DOES NOT COMPILE: cannot use f (type Foo) as type *Foo!

	// all in one line
	SetBar(&f, 42)
	fmt.Printf("After SetBar(&f, 42), Bar = %d\n", f.Bar)
}

// END OMIT
