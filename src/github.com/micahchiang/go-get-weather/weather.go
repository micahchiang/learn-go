package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello! Where did you want to check the weather?")
	fmt.Println("Type the name of the city, then press enter")
	scanner.Scan()
	city := scanner.Text()
	fmt.Println("Retrieving the weather for", city)
}
