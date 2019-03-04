package main

import (
	"net/http"
	"fmt"
	)

var urlstring = []string{
	"www.baidu.com",
	"www.taobao.com",
}
func main() {
	for _, v := range urlstring {
		res, err := http.Head(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Status)

	}
}