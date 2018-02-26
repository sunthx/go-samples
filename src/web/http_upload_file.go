package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func process(writer http.ResponseWriter,request *http.Request) {
	request.ParseMultipartForm(1024)
	fileHeader := request.MultipartForm.File["uploaded"][0]
	file,err := fileHeader.Open()
	if err == nil {
		data,err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Println(writer,string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
