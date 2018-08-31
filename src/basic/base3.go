package main

import "fmt"

func main() {
	i:=1;

	//影子变量
	if i:="123";len(i)> 0 {
		fmt.Println(i)
	}

	fmt.Println(i)
}
