package main

import (
	"net/http"
	"fmt"
)

func main()  {

	customHandler := CustomHandler{}
	//创建一个Http的服务
	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler: &customHandler,
	}

	//开启Http服务
	//server.ListenAndServe()

	//开始Https服务
	server.ListenAndServeTLS("cert.pem","key.pem")
}

type CustomHandler struct {

}

func (handler *CustomHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request)  {
	fmt.Fprint(writer,"Hello GoLang!")
}
