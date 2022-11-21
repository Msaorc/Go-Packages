package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/site", Site{name: "My Site"})
	http.ListenAndServe(":9000", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type Site struct {
	name string
}

func (s Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.name))
}
