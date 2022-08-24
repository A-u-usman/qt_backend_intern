package main

import "fmt"

//author: Muhammad Abdullahi

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for loop using range
	for _, num := range nums {
		fmt.Print(num)
	}
}
