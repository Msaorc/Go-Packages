package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.facebook.com")
	if err != nil {
		panic(err)
	}

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))
}
