package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// a pointer is an object that stores the memory address of another value located in memory. It's a reference
	// to that location. Obtaining the value at that memory location is known as dereferrencing the pointer.
	// Example: a page number in a book's index points to that physical page's location in the book.

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
	// notice above how calling zeroval doesn't change the value of i in main, but zeroptr does.
	// this is because zeroval gets passed a copy, while zeroptr is passed a reference.
}
