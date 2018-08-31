package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.List

	list.PushBack(1)
	list.PushBack(2)

	fmt.Println("succeed")
}
