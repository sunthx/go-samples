package main

import (
	"fmt"
)

func main() {
	// & 取值操作符
	// * 解引用操作符
	x:=37
	i:=x
	pi:= &x
	ppi:=&pi
	fmt.Println(x,i,pi,*pi,**ppi)

	**ppi++
	fmt.Println(x,i,pi,*pi,**ppi)

	fmt.Println("--------------")

	value1:=9
	value2:=5
	product:=0

	swapAndProduct1(&value1,&value2,&product)
	fmt.Println(value1,value2,product)

	swapAndProduct2(value1,value2)
	fmt.Println(value1,value2,product)

	fmt.Println("--------------")

	//类型值
	objectA := composer{"sunthInObjectA", 1990}

	//返回指向该类型值的指针
	objectB := new(composer)
	objectB.name= "sunthInObjectB"
	objectB.birthYear= 1991

	objectC := &composer{}
	objectC.birthYear=1992
	objectC.name= "sunthInObjectC"

	objectD := &composer{"sunthInObjectD", 1993}

	fmt.Println(objectA)
	fmt.Println(*objectB)
	fmt.Println(*objectC)
	fmt.Println(*objectD)


}

func swapAndProduct1(x,y,product *int){
	if *x > *y{
		*x,*y = *y,*x
	}
	*product = *x * *y
}

func swapAndProduct2(x,y int)(int,int,int){
	if x>y{
		x,y = y,x
	}

	return x,y,x*y
}

type composer struct {
	name 		string
	birthYear 	int
}