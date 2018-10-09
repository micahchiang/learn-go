package main

import "fmt"

func main() {
	// arrays have length and capacity
	var a [5]int
	fmt.Println(a)
	fmt.Println(cap(a))

	// arrays are zero indexed and values can be assigned using array[index] syntax
	a[4] = 55
	fmt.Println(a)
	fmt.Println(a[4])

	// length
	fmt.Println("length:", len(a))

	// array one liner for declaration and initialization
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// arrays are one dimensional but multi-dimensional data structures can be composed by them
	var twod [2][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)

	/*
		Slices are more common in Go than arrays. Their type is defined by the elements they contain.
		To create a slice with non-zero length, use built in make function
	*/

	s := make([]string, 3)
	fmt.Println(len(s))

	// getting and setting is the same as arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	// Unlike arrays, slices allow you to add values to the end of the sequence using append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)
	fmt.Println(len(s))

	// Slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// Slices can be sliced, similar to arrays in javascript.
	// The slice[low:high] syntax is low inclusive but high exclusive
	l := s[2:5]
	fmt.Println(l)

	// one liner for slices declaration and initialization
	t := []string{"one", "two", "three"}
	fmt.Println(t)

	// slices can create multi-dimensional data structures
	sliceTwoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		sliceTwoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			sliceTwoD[i][j] = i + j
		}
	}

	fmt.Println(sliceTwoD)
}
