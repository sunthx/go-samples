package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(chan1)
	})

	select {
	case _, ok := <-chan1:
		if !ok {
			fmt.Println("The candidate case is closed")
			break
		}
		fmt.Println("The candidate case is selected")
	}
}
