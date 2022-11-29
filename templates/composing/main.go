package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name   string
	Charge int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Courses{
			{"Go", 50},
			{"Java", 20},
			{"C#", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":9090", nil)
}
