package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}

func handler(writer http.ResponseWriter,request *http.Request) {
	fmt.Fprintf(writer,"Hello Go Web Program, %s!",request.URL.Path[1:])
}
