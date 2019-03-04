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

	v1, err1 := c.Do("hset","books","abc",100)
	if err1 != nil {
		fmt.Println("hset failed", err1)
		return
	}
	fmt.Println("HSet sucess will return:",v1)

	v2, err2 := redis.Int(c.Do("HGET","books","abc"))
	if err2 != nil {
		fmt.Println("hget failed", err2)
		return
	}
	fmt.Println("HGet sucess will return:", v2)


}

