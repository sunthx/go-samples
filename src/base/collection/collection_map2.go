package main

import "fmt"

func main()  {
	map1:=map[string]int{"name1":1,"name2": 2}
	fmt.Println(map1["name3"])
	fmt.Println(map1["name1"])

	value,exist := map1["name1"]
	fmt.Println(value,exist)
}
