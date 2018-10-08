package main

import "fmt"

func main() {
	// ways to declare variables

	// go infers variable type
	var a = "my name is a"
	fmt.Println(a)
	// one or more vars
	var b, c int = 1, 2
	fmt.Println(b, c)
	// omit types to create multiple variables with different types
	var d, e, f = "string value", 8, 8.0
	fmt.Println(d, e, f)
	// no initialization 'zeros' a variable
	var g int
	fmt.Println(g)
	// shorthand for declaring and initializing
	h := "hello world"
	fmt.Println(h)
}
