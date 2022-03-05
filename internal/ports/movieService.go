package ports

import domain "github.com/Tambarie/movie-api/internal/core/domain/movie"

type MovieService interface {
	SaveComments(comment *domain.Comment) (*domain.Comment, error)
	GetComments(movieId int) (*[]domain.Comment, error)
	CountComments(movieId int) (int64, error)
}

type RedisService interface {
	SetMovie(key string, value *[]domain.Movie) error
	GetMovie(key string) *[]domain.Movie
	SetMovieCharactersInRedis(key string, value []domain.Character) error
	GetMovieCharactersInRedis(key string) []domain.Character
}
