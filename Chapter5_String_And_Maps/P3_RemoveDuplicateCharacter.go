package main

import "fmt"

func main() {
	str := "golang"
	seen := make(map[rune]bool)
	result := ""
	for _, char := range str {
		if !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}
	fmt.Println("String after removing duplicates:", result)
}
