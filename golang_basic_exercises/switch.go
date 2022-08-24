package main

import "fmt"

func main() {
	a := 5
	switch a {
	case 1:
		fmt.Printf("The integer == %d", a)
		fallthrough
	case 2:
		fmt.Printf("The integer == %d", a)
	case 3:
		fmt.Printf("The integer == %d", a)
		fallthrough
	case 4:
		fmt.Printf("The integer == %d", a)
	case 5:
		fmt.Printf("The integer == %d", a)
		fallthrough
	default:
		fmt.Println("the integer did not satisfy all the switch cases")
	}
}
