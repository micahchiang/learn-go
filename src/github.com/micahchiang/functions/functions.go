package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hello,", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye,", name)
}

func beSociable(name string) {
	sayHello(name)
	fmt.Println("hows the weather?")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

// func sayHelloABunch() {
// 	fmt.Println("hello")
// 	sayHelloABunch()
// }

func main() {
	beSociable("bob")
	beSociable("alice")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))

	fmt.Println(x)
}
