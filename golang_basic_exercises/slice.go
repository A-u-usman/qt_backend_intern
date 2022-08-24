package main

import "fmt"

func main() {
	/* define a string */
	mystring := "Hello world"

	/* take slice */
	part_one := mystring[0:4]
	part_two := mystring[6:11]
	fmt.Printf("part one = %s \n", part_one)
	fmt.Printf("part two = %s \n", part_two)
}
