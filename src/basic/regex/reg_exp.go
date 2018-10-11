package main

import (
	"fmt"
	"regexp"
)

func main() {
	testData := "abc 123"
	testExp := `\w{3}|\d{3}`

	reg, _ := regexp.Compile(testExp)
	res := reg.FindAll([]byte(testData), -1)

	fmt.Println(string(res[0]))
	fmt.Println(string(res[1]))
}
