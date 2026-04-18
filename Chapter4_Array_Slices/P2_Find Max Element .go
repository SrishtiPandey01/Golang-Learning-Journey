package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	fmt.Println("Max element:", max)
}
