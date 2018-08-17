package main

import (
	"fmt"
)

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func addOne(num *int) {
	*num = *num + 1
}

func main() {
	x := 5
	fmt.Println(x)

	xPtr := &x // == var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p := position{4, 2}

	b := badGuy{"Jabba the Hut", 100, p}

	whereIsBadGuy(&b)
}
