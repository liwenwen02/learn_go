package main

import (
	"net/http"
	"time"
	"html/template"
	"fmt"
)
type Pagevariables struct {
	Date string
	Time string
}
func Homepage(rw http.ResponseWriter, r *http.Request) {
	timevar :=Pagevariables{
		Date: time.Now().Format("2006-01-02"),
		Time: time.Now().Format("15:04:05"),
	}
	t, err := template.ParseFiles("homepage.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(rw, timevar)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	http.HandleFunc("/", Homepage)

}
