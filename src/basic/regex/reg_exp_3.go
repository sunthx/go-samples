package main

import (
	"github.com/dlclark/regexp2"
	"fmt"
)

func main(){
	exp := `\w{3}|\d{3}`
	data := `123 456 abc #@!`

	res := regexp2.MustCompile(exp,0)
	if res == nil {
		fmt.Println("error")
		return
	}

	matchResult ,_ := res.FindStringMatch(data)
	if matchResult != nil {
		fmt.Println(matchResult)
	} else {
		fmt.Println("DONE")
		return
	}

	result1,_ := res.FindNextMatch(matchResult)
	if result1 != nil {
		fmt.Println(result1)
	} else {
		fmt.Println("DONE")
		return
	}

	fmt.Println("DONE!")

}
