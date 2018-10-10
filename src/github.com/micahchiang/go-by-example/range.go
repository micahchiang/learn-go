package main

import "fmt"

func main() {

	// range is used to iterate over a number data structures.

	// summing numbers in an array or slice:
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	// range provides both index and value at that index. If you don't need index, replace it with an underscore.

	// range on map iterates over key value pairs:
	m := map[string]string{"a": "apple", "b": "banana", "c": "cranberry"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range m {
		fmt.Printf("key: %s\n", k)
	}

	// range on strings iterates over unicode code points:
	for i, c := range "test" {
		fmt.Println(i, c)
	}
}
