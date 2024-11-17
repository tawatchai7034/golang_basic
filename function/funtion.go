package main

import "fmt"

func hello() {
	fmt.Println("hello 5555")
}

func calculater(num1 float64, num2 float64) float64 {
	result := num1 * num2
	fmt.Println(result)
	return result
}
func main() {
	hello()
	test := calculater(25.00, 80.3)
	fmt.Println("test", test)
}
