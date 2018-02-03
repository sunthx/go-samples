package main

import (
	"net/http"
	"fmt"
)

func main()  {
	defaultServeMux := http.DefaultServeMux

	httpServer:=http.Server{
		Addr:"127.0.0.1:8080",
		Handler:defaultServeMux,
	}

	handler1:=Handler1{}
	handler2:=Handler2{}

	http.Handle("/1",&handler1)
	http.Handle("/2",&handler2)

	handle4 := http.HandlerFunc(handleFunc4)

	//处理器函数
	http.HandleFunc("/3",handleFunc3)
	http.Handle("/4",&handle4)

	httpServer.ListenAndServe()
}

func handleFunc3(writer http.ResponseWriter,request *http.Request) {
	fmt.Fprint(writer,"3: handle func")
}

func handleFunc4(writer http.ResponseWriter,request *http.Request) {
	fmt.Fprint(writer,"4: handle func convert to handler")
}

type Handler1 struct {

}

func (handler *Handler1)ServeHTTP(writer http.ResponseWriter,request *http.Request)  {
	fmt.Fprint(writer,"1: default handler1")
}

type Handler2 struct {

}

func (handler *Handler2)ServeHTTP(writer http.ResponseWriter,request *http.Request)  {
	fmt.Fprint(writer,"2: default handler2")
}