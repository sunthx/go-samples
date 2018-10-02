package main

import (
	"fmt"
	"regexp"
)

func main() {
	test_data := "abc 123"
	test_exp := `\w{3}|\d{3}`

	reg, _ := regexp.Compile(test_exp)
	res := reg.FindAll([]byte(test_data), -1)

	fmt.Println(string(res[0]))
	fmt.Println(string(res[1]))
}
