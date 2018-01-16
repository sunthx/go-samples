package main

import (
	"fmt"
	"bookSamples/stackerples/stacker/stack"
)

func main()  {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"1"})
	haystack.Push(81.82)

	for{
		item,err:=haystack.Pop()
		if err != nil{
			break
		}
		fmt.Println(item)
	}

}