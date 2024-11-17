package main

import "fmt"

func switchCase(input string) string {
	grade := "F"
	switch input {
	case "A":
		grade = "A55"
		break
	case "B":
		grade = "B444"
		break
	case "C":
		grade = "C3333"
		break
	}
	return grade
}

func main() {
	condition := false
	if condition {
		fmt.Println(switchCase("A"))
	} else {
		fmt.Println(switchCase("B"))
	}
	fmt.Println(condition)
}
