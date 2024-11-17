package main

import "fmt"

var int1 int = 20
var list1 [3]string

func main() {
	list1[0] = "test"
	list1[1] = "retest"
	list1[2] = "55555"

	numfloat := 25.75

	priceList := [3]float64{10, 20, 30}
	fmt.Println(numfloat)
	fmt.Println("int1", int1)
	fmt.Println(list1)
	fmt.Println(priceList)
}
