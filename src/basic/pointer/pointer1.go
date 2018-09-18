package main

import (
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func New(name string) Dog {
	return Dog{name}
}

func main() {
	dog := New("dog1")
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
}
