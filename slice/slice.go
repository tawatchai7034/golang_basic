package main

import "fmt"

func main() {
	var list []string
	list = []string{"555", "444"}
	fmt.Println(list)
	list = append(list, "333", "222")
	fmt.Println(list)
	newList := list[2:4]
	fmt.Println(newList)
}
