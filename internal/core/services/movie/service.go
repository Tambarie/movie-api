package services

import (
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/ports"
)

type Service struct {
	movieRepository ports.MovieRepository
}

func (s Service) GetComments(movieId int) (*[]domain.Comment, error) {
	return s.movieRepository.GetComments(movieId)
}

func (s Service) CountComments(movieId int) (int64, error) {
	return s.movieRepository.CountComments(movieId)
}

func (s Service) SaveComments(comment *domain.Movie) (*domain.Movie, error) {
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

func (r *RedisService) SetMovieCharacters(key string, value []domain.Character) {
	return
}

func (r *RedisService) GetMovieCharacters(key string) []domain.Character {
	return r.redisRepository.GetMovieCharacters(key)
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
