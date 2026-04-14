package main
import "fmt"

func main() {
	var x int =10
	var y float64 = float64(x) // Type conversion from int to float64
	fmt.Println("Converted value of x to float64 is:", y)
}