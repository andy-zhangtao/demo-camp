package main

//直连Redis
import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

//addr 链接地址
func connectRedisWithDirect(addr string, passwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd, // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key1", "this is a demo string", time.Duration(1*time.Minute)).Err()
	if err != nil {
		panic(err)
	}

	result := client.Get("key1")
	fmt.Println(result.String())
}
