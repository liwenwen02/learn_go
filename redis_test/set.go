package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	c, err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("connection failed", err)
		return
	}
	defer c.Close()

	v1, err1 := c.Do("SET","name","li")
	if err1 != nil {
		fmt.Println("set failed", err1)
		return
	}
	fmt.Println("set sucess will return:",v1)

	v2, err2 := redis.String(c.Do("GET","name"))
	if err2 != nil {
		fmt.Println("get failed", err2)
		return
	}
	fmt.Println("get sucess will return:", v2)


}

