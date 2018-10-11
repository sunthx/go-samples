package main

import (
	"fmt"
)

func main() {
	goNum := 10
	sign := make(chan struct{})
	for i := 0; i < goNum; i++ {
		go func(index int) {
			defer func() {
				sign <- struct{}{}
			}()

			fmt.Printf(" current index : %d  \n",index)
		}(i)
	}

	for j := 0; j < goNum; j++ {
		<- sign
	}

	fmt.Println("DONE")
}
