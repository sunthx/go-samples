package main

import "time"

func main() {
	go func(){
		println("goroutine...")
		time.Sleep(time.Second * 5)
		println("goroutine...finished")
	}()

	time.Sleep(time.Second * 3)
	println("main exit!")
}
