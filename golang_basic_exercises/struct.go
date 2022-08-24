package main

import "fmt"

type House struct {
	noRooms int
	city    string
	price   int
}

func main() {
	var ahouse House

	ahouse.noRooms = 3
	ahouse.city = "Kaduna"
	ahouse.price = 30000000

	fmt.Printf("Number of rooms =  %d\n", ahouse.noRooms)
	fmt.Printf("Location  =  %s\n", ahouse.city)
	fmt.Printf("Price = %d", ahouse.price)
}
