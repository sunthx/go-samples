package main

func main() {
	chan1 := make(chan []string, 5)

	slice1 := make([]string, 5)
	slice1 = append(slice1, "demo")
	println(&slice1)

	chan1 <- slice1

	slice2 := <-chan1
	println(&slice2)
}
