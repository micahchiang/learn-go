package main

import "fmt"

// In Go, you can define methods on struct types.

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
