package main

import "fmt"
fun factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Factorial of 5 is:", factorial(5))
}
