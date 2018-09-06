package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channals := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("The index : %d\n", index)

	channals[index] <- index

	select {
	case <-channals[0]:
		fmt.Println("The first candidate case is selected.")
	case <-channals[1]:
		fmt.Println("The second candidate case is selected.")
	case third := <-channals[2]:
		fmt.Printf("The third candidate case is selected,the element is %d.\n", third)
	default:
		fmt.Println("No candidate case is selected!")
	}
}
