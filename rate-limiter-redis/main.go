package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	res, err := rdb.Eval(ctx, luaScript, []string{"test-key"}, "arg1").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
