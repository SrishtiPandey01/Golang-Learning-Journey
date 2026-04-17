package maqin

import "fmt"

func isPalindrome(num int) bool {
	originalNum := num
	reversedNum := 0
	for num > 0 {
		digit := num % 10
		reversedNum = reversedNum*10 + digit
		num /= 10
	}
	return originalNum == reversedNum
}

func main() {
	if isPalindrome(12321) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
