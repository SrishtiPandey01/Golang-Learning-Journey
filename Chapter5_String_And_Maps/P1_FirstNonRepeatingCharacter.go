package main

import "fmt"

func main() {
	str := "swiss"
	freq := make(map[rune]int)

	for _, char := range str {
		freq[char]++
	}

	for _, char := range str {
		if freq[char] == 1 {
			fmt.Println("First non-repeating character:", string(char))
			break
		}

	}
}
