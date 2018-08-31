package main

import (
	"math"
	"fmt"
)

func main()  {
	value,err := IntFromInt64(123213123123123211111123123)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(value)
	}
}

func IntFromInt64(x int64)(i int,err error) {
	defer func() {
		if e:= recover();e!=nil{
			err = fmt.Errorf("%v", e)
		}
	}()

	i= ConvertInt64ToInt(x)
	return i,nil;
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x<= math.MaxInt32{
		return int(x)
	}

	panic(fmt.Sprintf("%d is out of the int32 range", x))

}