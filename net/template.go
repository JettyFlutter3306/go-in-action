package net

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func TemplateReplace() {
	name := "luobida"
	myTemplate := "hello, {{.}}"
	tmpl, err := template.New("test").Parse(myTemplate)
	if err != nil {
		fmt.Print(err)
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println()

	person := Person{"洛必达", 18}
	myTemplateA := "hello, {{.Name}}, your age is {{.Age}}"
	tmpl, err = template.New("test").Parse(myTemplateA)
	if err != nil {
		fmt.Print(err)
	}
	err = tmpl.Execute(os.Stdout, person)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println()
}
