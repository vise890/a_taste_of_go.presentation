package main

import (
	"fmt"
)

// START OMIT
type Foo struct{ Bar int }

// f is passed by value, and Bar is updated in the copy
func UselessSetBar(f Foo, newBar int) { f.Bar = newBar }

// f is passed by reference:
//   - the argument `f` is of type *Foo (read pointer to Foo)
//   - this time bar will be updated in the Foo that we pass in
func SetBar(f *Foo, newBar int) { f.Bar = newBar }

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, Bar = %d\n", f.Bar)

	UselessSetBar(f, 5)                                  // when we pass by copy,
	fmt.Printf("After UselessSetBar, Bar = %d\n", f.Bar) // ..the change doesn't affect the original

	// SetBar(f, 5) // DOES NOT COMPILE: cannot use f (type Foo) as type *Foo!

	var aPointerTo_f *Foo = &f                    // we create a reference to f with `&` ("address")
	SetBar(aPointerTo_f, 5)                       // ..we pass it to SetBar..
	fmt.Printf("After SetBar, Bar = %d\n", f.Bar) // ..and the change happens on the original f

	SetBar(&f, 42) // or, all in one line
	fmt.Printf("After SetBar (second time), Bar = %d\n", f.Bar)
}

// END OMIT
