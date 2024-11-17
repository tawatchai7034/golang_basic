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
	e := employee{}
	data := `{"Id":12,"Name":"fluke","Tel":"0611804157","Email":"example@gmail.com"}`
	err := json.Unmarshal([]byte(data), &e)
	if err != nil {
		panic(err)
	}
	fmt.Println(e)
}
