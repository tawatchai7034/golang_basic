package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id    int
	Name  string
	Tel   string
	Email string
}

func main() {
	data, err := json.Marshal(&employee{12, "fluke", "0611804157", "example@gmail.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
