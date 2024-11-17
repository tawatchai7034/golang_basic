package main

import "fmt"

func zerovalue(i int) {
	i = 0
}

func zeropointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("init i = ", i)
	zerovalue(i)
	fmt.Println("i from zerovalue: ", i)
	zeropointer(&i)
	fmt.Println("i from zeropointer: ", i)
	fmt.Println("i from zeropointer address ", &i)
}
