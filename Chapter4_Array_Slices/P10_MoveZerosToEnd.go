package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	result := []int{}

	for _, v := range arr {
		if v != 0 {
			result = append(result, v)
		}

	}
	for len(result) < len(arr) {
		result = append(result, 0)
	}
	fmt.Println("Array after moving zeros to end:", result)
}
