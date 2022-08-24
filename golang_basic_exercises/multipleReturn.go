package main

import "fmt"

func values() (string, string) {
	return "hello", "world"
}

func main() {
	x, y := values()
	fmt.Println(x)
	fmt.Println(y)
}
