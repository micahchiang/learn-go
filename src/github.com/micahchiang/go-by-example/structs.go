package main

import "fmt"

// In Go, structs are typed collections of fields.
// They're useful for grouping data together to form records.

// The person struct below has name and age fields.
type person struct {
	name string
	age  int
}

func main() {
	// creating a new struct:
	p1 := person{"bob", 27}
	fmt.Println(p1)

	// you can also use the name of the fields during initialization:
	p2 := person{name: "alice", age: 27}
	fmt.Println(p2)

	// ommitted fields will be zero-valued:
	p3 := person{name: "jimbo"}
	fmt.Println(p3)

	// & prefix yields a pointer to the struct:
	fmt.Println(&person{name: "ann", age: 30})

	// struct fields are accessed with dot notation:
	s := person{name: "john", age: 25}
	fmt.Println(s.name)
	// using struct pointers automatically dereferrences:
	sp := &s
	fmt.Println(sp.age)
	// structs are mutable:
	sp.age = 48
	fmt.Println(s)
}
