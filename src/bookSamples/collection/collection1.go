package main

import "fmt"

func main() {
	array1:=[]string{"1","2", "3"}
	array2:=[]string{"4","5", "6"}
	array3:=[]string{"7","8", "9"}

	array1 = append(array1,"11","22", "33")
	fmt.Println(array1)

	array1 = append(array1,array2...)
	fmt.Println(array1)

	array1 = append(array1,array3[:2]...)
	fmt.Println(array1)
}
