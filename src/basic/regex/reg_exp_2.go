package main

import (
	"github.com/dlclark/regexp2"
	"fmt"
)

func main() {
	var testData = `<html><a href="http://www.baidu.com" target="_parent" /></html>`
	var exp = `(?<=<a.*href=").*(?="\s*target.*\/>)|(?<=<a.*\s*target=").*(?="\s?.*\/>)`

	regexp := regexp2.MustCompile(exp, 0)
	res, err := regexp.FindStringMatch(testData)
	if err != nil {
		fmt.Println("error")
		return
	}

	groups := res.Groups()
	fmt.Println(len(groups))
	fmt.Println(groups[0].Captures[0].String())

	fmt.Println(res)
}
