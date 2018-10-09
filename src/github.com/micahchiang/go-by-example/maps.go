package main

import "fmt"

func main() {

	// maps are Go's built in associative data type. In javascript these are objects. In python dicts.

	// use make to create an empty map: make(map[key]value-type)
	m := make(map[string]int)

	// key-value pairs are assigned in typical syntax:
	m["key1"] = 7
	m["key2"] = 8
	fmt.Println(m)

	// map values are retrived by their key name:
	v := m["key2"]
	fmt.Println(v)

	// len returns the number of key value pairs when used on a map:
	fmt.Println(len(m))

	// builtin `delete` function removes key value pairs:
	delete(m, "key1")
	fmt.Println(m)

	// second return values when retrieving key-values determine if key was/is present:
	_, prs := m["key1"]
	fmt.Println(prs)

	// one liner for declaring and intializing a map:
	n := map[string]int{"one": 1, "two": 2}
	fmt.Println(n)

}
