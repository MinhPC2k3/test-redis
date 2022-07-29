package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type account struct {
	name   string
	gender string
}

//var value=[]byte("value1")
func TestRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	minh := "congminh"
	check, _ := rdb.Exists(ctx, "key1").Result()
	if check == 0 {
		err := rdb.Set(ctx, "key1", minh, 0).Err()
		if err != nil {
			panic(err)
		}
	} else if check != 0 {
		fmt.Println("Key already exists")
	}
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)
	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	//     fmt.Println("key2 does not exist")
	// } else if err != nil {
	//     panic(err)
	// } else {
	//     fmt.Println("key2", val2)
	// }

}
func main() {
	TestRedis()
}
