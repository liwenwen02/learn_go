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

	_, err = c.Do("expire","name",5)
	if err != nil {
		fmt.Println("lpush failed", err)
		return
	}



}

