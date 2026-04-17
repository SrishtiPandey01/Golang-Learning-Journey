package main

import "fmt"

func arraySum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := arraySum(numbers)
	fmt.Println("Sum of the array:", result)
}
