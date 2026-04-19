package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}
