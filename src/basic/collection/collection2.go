package main

import "fmt"

func main() {
	array1:=[]string{"1","2","3","4", "5"}
	array2:=[]string{"1","2","3","4", "5"}

	fmt.Println(len(array1),cap(array1))

	x:= InsertStringSliceCopy(array1,[]string{"0"}, 0)
	x2:= InsertStringSlice(array2,[]string{"0"}, 0)

	fmt.Printf("%v\n",x)
	fmt.Printf("%v\n",x2)
	fmt.Printf("%v\n",array2)

	y:= InsertStringSliceCopy(array1,[]string{"6"},len(array1))
	y2:= InsertStringSlice(array2,[]string{"6"},len(array2))

	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", y2)
	fmt.Printf("%v\n",array2)

	z:=InsertStringSliceCopy(array1,[]string{"----"},3)
	z2:=InsertStringSlice(array2,[]string{"----"},3)

	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", z2)

	fmt.Println(array1)
	fmt.Println(array2)


}

func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string,len(slice)+len(insertion))
	at := copy(result,slice[:index])
	at += copy(result[at:],insertion)
	copy(result[at:],slice[index:])
	return result;
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index],append(insertion,slice[index:]...)...)
}