package main

import (
	"bytes"
	"sync"
	"io"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	var buffer bytes.Buffer
	var mutex sync.Mutex
	signal := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func(id int,writer io.Writer){
			defer func(){
				signal <- struct{}{}
			}()

			for j := 0; j < 10; j++ {
				data := fmt.Sprintf("\n[id: %d, iteration: %d]",id,j)
				mutex.Lock()
				_,err := writer.Write([]byte(data))
				if err != nil {
					log.Printf("error: %s [%d]",err,id)
				}
				mutex.Unlock()
			}
		}(i,&buffer)
	}

	for i := 0; i < 10; i++ {
		<- signal
	}

	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	log.Printf("The contents:\n%s", data)
}
