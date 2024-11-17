package main

import "fmt"

func add(result1, result2 float64) float64 {
	return result1 + result2
}

func loop(text string) {
	for i := 0; i < 100; i++ {
		fmt.Println(text, " : ", i)
	}
}

func main() {
	defer fmt.Println("test1")
	fmt.Println("55555")
	defer loop("defer")
	fmt.Println(add(5, 7.25))
}
