package channel

import (
	"fmt"
	"time"
)

func loop(text string) {
	for i := 0; i < 100; i++ {
		fmt.Println(text, " : ", i)
	}
}

func main() {
	loop("test1")
	go loop("loop")
	time.Sleep(5 * time.Second)
}
