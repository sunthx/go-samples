package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(writer http.ResponseWriter,request *http.Request,param httprouter.Params)  {
	fmt.Fprintf(writer,"hello,%s!\n",param.ByName("name"))
}

func main()  {
	mux := httprouter.New()
	mux.GET("/hello/:name",hello)

	server := http.Server {
		Addr: "127.0.0.1:8080",
		Handler:mux,
	}

	//server.ListenAndServe()

	server.ListenAndServeTLS("cert.pem","key.pem")
}
