package main

import (
	"fmt"
)

func sum(a, b int) int {
	//call add to add the numbers
	//function calling another function
	result := add(a, b)
	return result
}

func add(a, b int) int {
	result := a + b
	return result
}
func main() {
	// call sum to add two number
	total := sum(10, 20)
	fmt.Println(total)

	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
}
