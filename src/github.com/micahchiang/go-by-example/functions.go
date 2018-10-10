package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// Go supports multiple return values from functions:
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// you can also return a subset of the returned values:
	_, c := vals()
	fmt.Println(c)

	// Go supports variadic functions with any number of parameters:
	nums := []int{1, 2, 3, 4, 5}
	sums(nums...)
	nums2 := []int{4, 5, 6, 7, 2, 4, 65, 67, 2}
	sums(nums2...)

	// Go supports closures:
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Go supports recursion:
	fmt.Println(fact(7))
}
