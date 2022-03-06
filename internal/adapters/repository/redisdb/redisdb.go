package redisdb

import (
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisClient(host string, db int, expiry time.Duration) ports.RedisRepository {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: expiry,
	}

}

func (r *RedisCache) getClient() *redis.Client {
	var redisURL *redis.Options
	if os.Getenv("REDIS_URL ") == "" {
		redisURL = &redis.Options{
			Addr:     r.host,
			Password: "",
			DB:       r.db,
		}
	} else {
		var err error
		redisURL, err = redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}

	}
	return redis.NewClient(redisURL)
}
