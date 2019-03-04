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

	_, err1 := c.Do("MSet","abc",100,"def",200)
	if err1 != nil {
		fmt.Println("mset failed", err1)
		return
	}


	v, err2 := redis.Ints(c.Do("MGET","abc","efg")) //此处是Ints
	if err2 != nil {
		fmt.Println("mget failed", err2)
		return
	}

	for _,val := range v {
		fmt.Println(val)
	}
	//未找到值时返回0


}

