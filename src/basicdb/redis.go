package basicdb

import "github.com/redis/go-redis/v9"

type RedisDatabase struct {
	Id string
	Instance *redis.Client
}

func (c *RedisDatabase) Connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	c.Instance = rdb
}