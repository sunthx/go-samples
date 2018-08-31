package main

import "fmt"

func main() {
	classifier(123123213.22, 12123, "string", nil, false)
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("bool,%d", i)
		case float64:
			fmt.Printf("float64,%d", i)
		case int, int32, int16, int64:
			fmt.Printf("int,%d", i)
		case nil:
			fmt.Printf("nil,%d", i)
		case string:
			fmt.Printf("string,%d", i)
		default:
			fmt.Printf("default,%d", i)
		}
	}

}
