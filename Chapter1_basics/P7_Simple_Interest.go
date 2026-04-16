package main

import "fmt"

func main() {
	p := 1000.0
	r := 5.0
	t := 2.0

	si := (p * r * t) / 100

	fmt.Println("Simple Interest:", si)
}
