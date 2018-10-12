package main

import (
	"fmt"
	"runtime"
)

//runtime.Gosched
//释放当前线程，调度执行别的任务
func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {
			fmt.Println("b")
		}()

		for i := 0; i < 4; i++ {
			fmt.Println("a", i)

			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	<-exit

	fmt.Println("DONE")
}
