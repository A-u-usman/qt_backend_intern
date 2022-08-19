package main

import "fmt"

func main() {

	num_array := [...]int{1, 2, 3, 4, 5}
	info_array := [...]string{"Muhammad Abdullahi", "backend", "Golang"}
	for i, num := range num_array {
		fmt.Printf(" num_array[%d] = ", i)
		fmt.Println(num)
	}
	for i, info := range info_array {
		fmt.Printf("Info_array[%d] = ", i)
		fmt.Println(info)
	}

}
