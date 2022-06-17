package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	err := initialize()
	if err != nil {
		return 
	}
}

func initialize() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		PoolSize: 100,
	})
	
	_, err = rdb.Ping(ctx).Result()
	return
}

func TestConnect(t *testing.T) {
	if err := initialize(); err != nil {
		fmt.Printf("Connected Failed Err: %v\n", err)
		panic(err)
	}
	
	fmt.Println("Connect Successfully!")
}

func TestSetString(t *testing.T) {
	err := rdb.Set(ctx, "name", "value1", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestGetString(t *testing.T) {
	// Get(ctx, key).Val()
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)
}

func TestGetHash(t *testing.T) {
	v := rdb.HGetAll(ctx, "user").Val()
	fmt.Println(v)
	
	v2 := rdb.HMGet(ctx, "user", "name", "age").Val()
	fmt.Println(v2)
}




