package main

import (
	"net/http"
	"fmt"
	"strings"
)
func GetFormInfo(rw http.ResponseWriter, r *http.Request){
	r.ParseForm()
	for _,v := range r.Form {
		fmt.Println(strings.Join(v,""))
	}
	fmt.Fprint(rw)

}
func main(){

	http.HandleFunc("/",GetFormInfo)
}
