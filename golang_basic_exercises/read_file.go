package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// red line by line into memory
// all file content is store in lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	//open file for reading
	//read line by line
	lines, err := readLines("names.txt")
	if err != nil {
		log.Fatalf("readLine :%s", err)
	}
	fmt.Println(lines[0])
	for i, name := range lines {
		fmt.Printf("line %d :: %s \n", i, name)
	}

}
