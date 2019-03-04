package main

import (
	"html/template"
	"os"
	"fmt"
)

const introduction = `my name is {{.Name}}, my age is {{.Age}}`
type PersonInfo struct {
	Name string
	Age int
}
func main() {
	li_info := PersonInfo{"li",25}
	t := template.Must(template.New("intro").Parse(introduction))
	err := t.Execute(os.Stdout, li_info)
	if err != nil {
		fmt.Println(err)
	}
}