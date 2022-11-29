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
	course := Course{"Templete Must", 50}
	t := template.Must(template.New("CourseTempleteMust").Parse("Course: {{.Name}} - Workload: {{.Charge}}"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
