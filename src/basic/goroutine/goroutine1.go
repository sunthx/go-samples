package main

import "fmt"

func main()  {
	counterA := createCounter(1)
	counterB := createCounter(101)

	for i := 0; i < 10; i++ {
		a := <- counterA
		b := <- counterB

		fmt.Printf("{a:%d,b:%d}\n",a,b)

	}
}

func createCounter(start int) chan int{
	next := make(chan int)
	go func(i int){
		for{
			next <- i
			i++
		}
	}(start)
	return next;
}