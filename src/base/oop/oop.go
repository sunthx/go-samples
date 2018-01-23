package main
import (
	"fmt"
	"reflect"
)

//实现泛型函数
func main()  {
	i:=Minimum(1,2,3,4,5,6).(int)
	fmt.Printf("%T %v\n", i,i)

	xs:=[]int{2,4,5,6,7}
	fmt.Println("6 @",Index(xs,6))
	fmt.Println("2 @",IndexReflectX(xs,2))

	fmt.Println("7 @",IndexReflectX2(xs,7))
}

func Minimum(first interface{},rest...interface{})interface{} {
	minimum := first
	for _,x := range rest{
		switch x:=x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}

		case float64:
			if x<minimum.(float64) {
				minimum = x
			}
		case string:
			if x<minimum.(string) {
				minimum = x
			}
		}

	}

	return minimum
}

//处理切片类型
func Index(xs interface{},x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i,y := range slice{
			if y == x.(int) {
				return i
			}
		}
	case []string:
		for i,y := range slice{
			if y==x.(string) {
				return i
			}
		}

	}
	return -1
}

func IndexReflectX(xs interface{},x interface{}) int {
	if slice := reflect.ValueOf(xs);slice.Kind()==reflect.Slice{
		for i:=0;i<slice.Len();i++{
			switch y:=slice.Index(i).Interface().(type) {
			case int:
				if y==x.(int) {
					return i
				}
			case string:
				if y==x.(string) {
					return i
				}

			}
		}
	}
	return -1
}

func IndexReflectX2(xs interface{},x interface{}) int{
	if slice := reflect.ValueOf(xs);slice.Kind()==reflect.Slice{
		for i:=0;i<slice.Len();i++{
			if reflect.DeepEqual(x,slice.Index(i)) {
				return i
			}
		}
	}
	return -1
}