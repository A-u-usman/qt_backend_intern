package main

import "fmt"

func names(name ...string) {

	for _, eachname := range name {

		fmt.Println(eachname)
	}

}

func main() {
	names("musa", "sale", "bintu")
}
