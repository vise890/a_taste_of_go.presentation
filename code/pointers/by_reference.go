package main

import (
	"fmt"
)

// START OMIT
type Foo struct{ Bar int }

// f is passed by reference:
//   - the argument `f` is of type *Foo (read: pointer to Foo)
//   - this time bar will be updated in the Foo that we pass in
func (f *Foo) SetBar(newBar int) { f.Bar = newBar }

func main() {
	f := Foo{Bar: 1}
	fmt.Printf("Originally, f.Bar = %d\n", f.Bar)

	// we create a reference to f with `&` ("address")
	var aPointerTo_f *Foo = &f
	// ..we pass it to SetBar..
	aPointerTo_f.SetBar(5)

	// ..and the change happens on the original f
	fmt.Printf("After aPointerTo_f.SetBar(5), f.Bar = %d\n", f.Bar)

	// NOTE: you can just call `func`s on pointers..
}

// END OMIT
