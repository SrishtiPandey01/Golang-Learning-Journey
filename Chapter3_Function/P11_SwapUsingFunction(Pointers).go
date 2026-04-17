package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println("After swap:", x, y)
}
