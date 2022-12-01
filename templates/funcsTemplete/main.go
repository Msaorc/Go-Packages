package main

import (
	"os"
	"strings"
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
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 50},
		{"Java", 20},
		{"C#", 10},
	})
	if err != nil {
		panic(err)
	}
	// })
	// http.ListenAndServe(":9090", nil)
}
