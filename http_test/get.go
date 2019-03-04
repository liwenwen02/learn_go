package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func main() {
	res, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
