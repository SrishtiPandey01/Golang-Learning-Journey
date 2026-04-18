package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	key := 30
	found := false

	for _, v := range arr {
		if v == key {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Element found in the array.")
	} else {
		fmt.Println("Element not found in the array.")
	}
}
