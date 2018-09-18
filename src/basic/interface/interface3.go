package main

import (
	"fmt"
)

type Pet interface {
	Name() string
}

type Dog struct {
	name string
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	var dog1 *Dog
	if dog1 == nil {
		fmt.Println("dog1 is nil")
	}

	dog2 := dog1
	if dog2 == nil {
		fmt.Println("dog2 is nil")
	}

	var pet Pet = dog2
	if pet == nil {
		fmt.Println("pet is nil")
	} else {
		fmt.Println("pet is not nil")
	}
}
