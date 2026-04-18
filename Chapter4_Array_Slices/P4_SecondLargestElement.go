package main

import "fmt"

func main() {
	arr := []int{10, 20, 5, 30, 15}
	first, second := 0, 0

	for _, v := range arr {
		if v > first {
			second = first
			first = v
		} else if v > second && v != first {
			second = v
		}
	}
	fmt.Println("Second largest element:", second)
}
