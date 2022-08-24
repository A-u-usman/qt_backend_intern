package main

import (
	"fmt"
	"os"
)

// Author: Muhammad Abdullahi
func main() {

	if _, err := os.Stat("for_loop.go"); os.IsNotExist(err) {
		fmt.Printf("file does not  Exixts\n")
	} else {
		fmt.Println("file exist!!!")
	}
}
