package main

import (
	"fmt"
	"os"
)

// author : Muhammad Abdullahi
func main() {
	file, err := os.Create("cities.txt")

	if err != nil {
		return
	}
	defer file.Close()
	cities := []string{"kaduna", "Abuja", "borno", "jigawa"}

	for i := 0; i < len(cities); i++ {
		file.WriteString(cities[i] + "\n")

	}
	fmt.Println("cities successfully added to the file cities.txt")
}
