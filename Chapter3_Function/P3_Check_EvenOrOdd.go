package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	if isEven(10) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
