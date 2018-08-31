package main

import "fmt"

func main() {
	xs := []int{2,3,4,5, 6}
	x:=IntSlice(xs)
	fmt.Println("2 @",IndexSlicer(x,3))
}

func IndexSlicer(slice slicer,x interface{}) int {
	for i:=0;i<slice.Len();i++{
		if slice.EqualTo(i,x) {
			return i
		}
	}
	return -1;
}

type slicer interface {
	EqualTo(i int,x interface{}) bool
	Len() int
}

type IntSlice []int
func (slice IntSlice) EqualTo(i int,x interface{}) bool{
	return slice[i] == x.(int)
}

func (slice IntSlice) Len() int  {
	return len(slice)
}
