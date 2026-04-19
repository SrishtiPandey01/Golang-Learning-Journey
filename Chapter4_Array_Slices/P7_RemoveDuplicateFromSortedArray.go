package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 4}
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	fmt.Println("Array after removing duplicates:", result)

}
