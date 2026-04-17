package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if isPrime(18) {
		fmt.Println("Prime")
	} else {

		fmt.Println("Not Prime")
	}
}
