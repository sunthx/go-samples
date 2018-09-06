package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int, 3)
	chan1 <- 1
	chan1 <- 2
	chan1 <- 3

	value := <-chan1

	fmt.Println(value)
}
