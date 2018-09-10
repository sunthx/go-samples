package main

import (
	"fmt"
)

type Person struct {
	name string
	age  string
}

type Teacher struct {
	school string
	Person
}

func (t Teacher) String() string {
	return fmt.Sprintf("%s school:%s", t.Person, t.school)
}

func (p Person) String() string {
	return fmt.Sprintf("name: %s,age: %s", p.name, p.age)
}

func main() {
	person := Person{name: "sunth", age: "28"}
	fmt.Println(person)

	teacher := Teacher{Person: person, school: "bd"}
	fmt.Println(teacher)
}
