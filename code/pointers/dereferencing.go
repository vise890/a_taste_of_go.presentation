package main

import (
	"fmt"
)

// START OMIT
type Foo struct{ Bar int }

func (f *Foo) SetBar(newBar int) { f.Bar = newBar }

func main() {
	f := Foo{Bar: 1}
	var pointerTo_f *Foo = &f

	// dereferencing pointers

	// just put an asterisk in front of a pointer to get what it points to:
	fmt.Println("*pointerTo_f = f is", *pointerTo_f == f)

	newF := Foo{Bar: 123}
	*pointerTo_f = newF
	// pointerTo_f now points to `newF`:
	fmt.Printf("pointerTo_f now points to ", *pointerTo_f)

}

// END OMIT
