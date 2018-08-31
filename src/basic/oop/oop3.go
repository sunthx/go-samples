package main

import (
	"strings"
	"fmt"
)

type Part struct {
	Id   int
	Name string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)

}
func main() {
	part := Part{1, "sunth"}
	part.LowerCase()
	part.Id += 1
	fmt.Println(part, part.HasPrefix("s"))
}
