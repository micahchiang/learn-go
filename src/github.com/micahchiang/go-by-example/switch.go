package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statements across multiple branches:

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// multiple expressions can be used in the same case:
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("It's a weekday. Sad!")
	}

	// switch with no expression is another way to express if/else logic:
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// switches can compare types:
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("I don't know what type %T is \n", t)
		}
	}
	whatAmI(true)
	whatAmI(8)
	whatAmI("hello")
	whatAmI(7.4)
}
