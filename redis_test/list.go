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

	_, err1 := c.Do("lpush","book_list","abc","def")
	if err1 != nil {
		fmt.Println("lpush failed", err1)
		return
	}


	v, err2 := redis.String(c.Do("lpop","book_list"))
	if err2 != nil {
		fmt.Println("lpop failed", err2)
		return
	}
	fmt.Println("lpop sucess will return:", v)


}

