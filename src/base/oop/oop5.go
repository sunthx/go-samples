package main

import "fmt"

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first,second string
}

func (pair *StringPair) Exchange()  {
	pair.first,pair.second = pair.second,pair.first
}

type Point [2]int

func (point *Point) Exchange()  {
	point[0],point[1] = point[1],point[0]
}

func exchangeThere(exchangers...Exchanger)  {
	for _,exchanger := range exchangers{
		exchanger.Exchange()
	}

}

func main() {
	value1 := StringPair{"first", "second"}
	value2 :=StringPair{"first2", "second2"}
	value3 :=Point{1, 2}

	fmt.Println("before:",value1,value2, value3)

	value1.Exchange()
	value2.Exchange()
	value3.Exchange()

	fmt.Println("after:",value1,value2, value3)

	exchangeThere(&value1,&value2,&value3)
	fmt.Println("after #2:",value1,value2,value3)
}