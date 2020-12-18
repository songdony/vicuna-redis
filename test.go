package main

import (
	"fmt"
	"time"

	"github.com/songdony/vicuna-redis/go_redis"
)

func main()  {
	//iter:=go_redis.NewStringOperation().MGet("name","name2").Iter()
	//
	//for iter.HasNext(){
	//	fmt.Println(iter.Next())
	//}

	//fmt.Println(go_redis.NewStringOperation().Get("name").Unwrap_Or("nothing"))

	fmt.Println(
		go_redis.NewStringOperation().
			Set("name","dony",
				go_redis.WithExpire(time.Second*25),
				go_redis.WithXX(),  //
				))
}