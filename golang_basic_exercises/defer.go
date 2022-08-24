package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("hello.txt")
	defer f.Close()

	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
}
