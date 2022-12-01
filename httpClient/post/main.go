package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	route := "http://endpointapi.com.br"
	json := bytes.NewBuffer([]byte(`{"name": "Marcos Augusto"}`))
	resp, err := c.Post(route, "application/json", json)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
