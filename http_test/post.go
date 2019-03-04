package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main() {
	res, err  := http.Post("https://www.baidu.com","application/x-www-form-urlencode",strings.NewReader("name=123"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}