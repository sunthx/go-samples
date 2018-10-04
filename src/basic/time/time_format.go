package main

import (
	"fmt"
	"time"
)

func main() {
	//format layout: 2006-01-02 15:04:05.999999999 -0700 MST
	cur := time.Now()
	fmt.Println(cur.Format("2006-01-02"))
}
