package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	clientRequest := http.Client{}
	route := "http://google.com"
	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")
	response, err := clientRequest.Do(request)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
