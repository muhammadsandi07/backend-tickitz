package pkg

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	redisHost := os.Getenv("RDSHOST")
	redisPORT := os.Getenv("RDSPORT")
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPORT),
	})
}
