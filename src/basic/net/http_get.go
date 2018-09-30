package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url_trending := "https://github.com/trending"
	resp, err := http.Get(url_trending)
	if err != nil {
		fmt.Println("request error")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request error")
		return
	}

	html := string(body)
	fmt.Println(html)
}
