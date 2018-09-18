package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

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
	dog := Dog{"little pig"}
	var pet Pet = dog
	dog.SetName("monster")

	//pet.Name() : little pig
	fmt.Println(pet.Name())

	//dog.Name() : monster
	fmt.Println(dog.Name())
}
