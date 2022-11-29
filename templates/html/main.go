package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name   string
	Charge int
}

type Courses []Course

func main() {
	t := template.Must(template.New("templete.html").ParseFiles("templete.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 50},
		{"Java", 20},
		{"C#", 10},
	})
	if err != nil {
		panic(err)
	}
}
