package main

import "fmt"

func main() {
	num := 15

	switch {
	case num%3 == 0 && num%5 == 0:
		fmt.Println("Divisible by both 3 and 5")
	case num%3 == 0:
		fmt.Println("Divisible by 3")
	case num%5 == 0:
		fmt.Println("Divisible by 5")
	default:
		fmt.Println("Not divisible by 3 or 5")
	}
}
