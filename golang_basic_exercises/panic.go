package main

import "os"

func main() {
	_, err := os.Create("/root/example")
	if err != nil {
		panic("Cannot create file")
	}
}
