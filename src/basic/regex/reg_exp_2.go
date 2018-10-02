package main

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

func main() {
	var test_data = `<html><a href="http://www.baidu.com" target="_parent" /></html>`
	var exp = `(?<=<a.*href=").*(?="\s*target.*\/>)|(?<=<a.*\s*target=").*(?="\s?.*\/>)`
	regexp := regexp2.MustCompile(exp, 0)
	res, err := regexp.FindStringMatch(test_data)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(res)
}
