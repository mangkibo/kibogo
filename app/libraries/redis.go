package libraries

import "github.com/go-redis/redis"

func Redis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     Env("REDIS_HOST") + ":" + Env("REDIS_PORT"),
		Password: Env("REDIS_PASSWORD"),
		DB:       0,
	})

	return redis
}
