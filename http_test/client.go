package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main(){
	client := &http.Client{}
	req, err := http.NewRequest("post","www.baidu.com",strings.NewReader("name = abc"))
	//strings.NewReader()创建Reader，并将内容
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
