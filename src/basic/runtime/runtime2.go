package main

import (
		"fmt"
	"runtime"
)

func main(){
	for i := 0; i<5;i++  {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}

	runtime.Goexit()

	fmt.Println("DONE")
}