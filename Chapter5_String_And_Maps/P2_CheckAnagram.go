package main

import "fmt"

func main() {
	str1 := "listen"
	str2 := "silent"

	if len(str1) != len(str2) {
		fmt.Println("Not anagrams")
		return
	}

	m := make(map[rune]int)

	for _, char := range str1 {
		m[char]++
	}

	for _, char := range str2 {
		m[char]--
	}

	for _, v := range m {
		if v != 0 {
			fmt.Println("Not anagrams")
			return
		}
	}

	fmt.Println("Anagrams")
}
