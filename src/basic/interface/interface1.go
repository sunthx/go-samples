package main

import (
	"fmt"
)

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

//Dog类型包含两个方法
//*Dog类型则包含三个方法
type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "DOG"
}

func main() {
	dog := Dog{"Toolcat"}
	dog.SetName("Toolcat2")
	fmt.Println(dog.Name())
	fmt.Println()
}
