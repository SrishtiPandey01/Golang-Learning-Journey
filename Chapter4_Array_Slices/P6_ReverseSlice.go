package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	fmt.Println("Reversed slice:", slice) // Output: [5 4 3 2 1]
}
