package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name   string
	Charge int
}

func main() {
	course := Course{"Go", 120}
	tmp := template.New("CourseTempleteGo")
	tmp, _ = tmp.Parse("Course: {{.Name}} - Workload: {{.Charge}}")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
