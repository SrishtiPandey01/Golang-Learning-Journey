package main

import "fmt"

func main() {
	str := "success"
	freq := make(map[rune]int)

	for _, char := range str {
		freq[char]++
	}

	maxCount := 0
	var result rune

	for k, v := range freq {
		if v > maxCount {
			maxCount = v
			result = k
		}
	}
	fmt.Println("Most frequent character:", string(result))
}
