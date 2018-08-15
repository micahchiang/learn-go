package main

import (
	"fmt"
)

func main() {
	// infinite loop
	// for {
	// 	fmt.Println("hello world")
	// }

	// for i < 10 {
	// 	fmt.Println("hello world", i)
	// 	i = i + 1
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	x := 5
	if x > 5 {
		fmt.Println("hello world")
	} else if x < 5 {
		fmt.Println("goodbye")
	} else {
		fmt.Println("it is five")
	}
}

/*
		Common Types

		var a int8 // -123.123
		var b uint8 // 0 - 255
		vac c int16
		var d uint16
		var e int // 32 bit on 32 bit machines 64 bit on 64bit machines

		var f1 float32
		var f2 float64

	// a := 2.0
	// b := 3.0
	// r := a / b
	// fmt.Printf("%.20f\n", r)

	// torf := true // implemented in int8

	/* Booleans
	var b bool = true
*/

/* strings
helloWorld := "Hellow World"
*/
