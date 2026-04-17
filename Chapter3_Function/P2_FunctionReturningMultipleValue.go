package main

import "fmt"

func calculate(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, difference := calculate(10, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}

//
//func main() {
//sum, difference := calculate(10, 5)
// fmt.Println("Sum:", sum, "Difference:", difference)

//}
