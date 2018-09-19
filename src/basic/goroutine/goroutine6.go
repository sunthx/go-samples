package main

import (
	"fmt"
	"sync/atomic"
)

//利用原子操作按照顺序打印 0 - 9
func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}

			trigger(i, fn)
		}(i)
	}

	trigger(10, func() {})
}
