package redisdb

import (
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/go-redis/redis/v8"
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
	return redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: "",
		DB:       r.db,
	})
}
