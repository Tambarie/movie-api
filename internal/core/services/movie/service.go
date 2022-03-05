package services

import (
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/ports"
)

type Service struct {
	movieRepository ports.MovieRepository
}

func (s Service) GetComments(movieID int) (*[]domain.Comment, error) {
	return s.movieRepository.GetComments(movieID)
}

func (s Service) CountComments(movieID int) (int64, error) {
	return s.movieRepository.CountComments(movieID)
}

func (s Service) SaveComments(comment *domain.Comment) (*domain.Comment, error) {
	return s.movieRepository.SaveComments(comment)
}

type RedisService struct {
	redisRepository ports.RedisRepository
}

func (r *RedisService) SetMovie(key string, value *[]domain.Movie) error {
	return r.redisRepository.SetMovie(key, value)
}

func (r *RedisService) GetMovie(key string) *[]domain.Movie {
	return r.redisRepository.GetMovie(key)
}

func (r *RedisService) SetMovieCharactersInRedis(key string, value []domain.Character) error {
	return r.redisRepository.SetMovieCharactersInRedis(key, value)
}

func (r *RedisService) GetMovieCharactersInRedis(key string) []domain.Character {
	return r.redisRepository.GetMovieCharactersInRedis(key)
}

func New(movieRepository ports.MovieRepository) *Service {
	return &Service{
		movieRepository: movieRepository,
	}
}

func NewRedisService(redisRepository ports.RedisRepository) *RedisService {
	return &RedisService{
		redisRepository: redisRepository,
	}
}
