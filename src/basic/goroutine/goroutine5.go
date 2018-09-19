package main

import (
	"fmt"
)

//使用chan等待子routine结束
func main() {
	fmt.Println("START")
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}

	for j := 0; j < num; j++ {
		<-sign
	}

	fmt.Println("DONE")
}
