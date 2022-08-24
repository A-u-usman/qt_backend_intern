package main

//Author: Muhammad Abdullahi
import (
	"fmt"
)

func main() {
	x := 6
	if x > 0 {
		fmt.Println(x / 2)
	} else {
		fmt.Printf("Cannot devide %d by 2 ", x)
	}

	i := 1
	max := 20

	// technically go doesnt have while, but
	// for can be used while in go.
	for i < max {
		fmt.Println(i)
		i += 1
	}

}
