package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2, 4, 2, 5}
	key := 2
	count := 0

	for _, value := range arr {
		if value == key {
			count++
		}
	}

	fmt.Println("Count:", count)
}
