package main

import "fmt"

var products = make(map[string]float64)

func main() {
	// Add
	products["A1"] = 5000
	products["A2"] = 2000
	products["A3"] = 1000
	products["A4"] = 6000
	fmt.Println("products = ", products)
	// Delete
	delete(products, "A1")
	fmt.Println("products = ", products)
	newMap := map[string]string{"B1": "FFF", "B2": "BBB"}
	fmt.Println(newMap)
}
