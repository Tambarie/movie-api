package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"log"
	"time"
)

func (r *RedisCache) SetMovie(key string, value *[]domain.Movie) error {
	log.Println("setting context")
	fmt.Println("setting context")
	client := r.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(context.Background(), key, string(json), r.expires*time.Minute).Err()
	if err != nil {
		return err
	}
	pong, err := client.Ping(context.Background()).Result()
	log.Println(pong, err)
	return err
}

func (r *RedisCache) GetMovie(key string) *[]domain.Movie {
	redisClient := r.getClient()
	val, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}
	var movie []domain.Movie
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}
	log.Println("Movies retrieved from cache")
	return &movie
}

func (r *RedisCache) SetMovieCharacters(key string, value []domain.Character) {
	return
}

func (r *RedisCache) GetMovieCharacters(key string) []domain.Character {
	return nil
}
