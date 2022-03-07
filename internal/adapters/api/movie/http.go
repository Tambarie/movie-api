package api

import (
	"github.com/Tambarie/movie-api/internal/ports"
)

type HTTPHandler struct {
	movieService ports.MovieService
	redisService ports.RedisService
}

func NewHTTPHandler(movieService ports.MovieService, redisService ports.RedisRepository) *HTTPHandler {
	return &HTTPHandler{
		movieService: movieService,
		redisService: redisService,
	}
}
