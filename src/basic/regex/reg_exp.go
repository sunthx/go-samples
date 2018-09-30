package main

import (
	"fmt"
	"regexp"
)

func main() {
	exp := "/b/w+(?=ing/b)"
	reg, err := regexp.CompilePOSIX(exp)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := reg.Match([]byte("sunfei"))
	if result {
		fmt.Println("Succeed")
	} else {
		fmt.Println("Error")
	}

	fmt.Println("DONE")
}
