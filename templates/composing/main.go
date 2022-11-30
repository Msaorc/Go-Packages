package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name   string
	Charge int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
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
