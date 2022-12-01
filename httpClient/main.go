package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	request := http.Client{Timeout: time.Second}
	req, err := request.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	response, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(response))
}
