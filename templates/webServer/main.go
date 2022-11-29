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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("templete.html").ParseFiles("templete.html"))
		err := t.Execute(w, Courses{
			{"Go", 50},
			{"Java", 20},
			{"C#", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":9000", nil)
}
