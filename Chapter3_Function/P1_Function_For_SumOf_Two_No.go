package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	result := sum(5, 10)
	fmt.Println("Sum:", result)
}
