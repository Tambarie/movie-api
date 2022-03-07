package ports

import domain "github.com/Tambarie/movie-api/internal/core/domain/movie"

// MovieRepository interface that connects the postgreDB
type MovieRepository interface {
	SaveComments(comment *domain.Comment) (*domain.Comment, error)
	GetComments(movieId int) (*[]domain.Comment, error)
	CountComments(movieId int) (int64, error)
}

// RedisRepository interface that connects the redisDB
type RedisRepository interface {
	SetMovie(key string, value *[]domain.Movie) error
	GetMovie(key string) *[]domain.Movie
	SetMovieCharactersInRedis(key string, value []domain.Character) error
	GetMovieCharactersInRedis(key string) []domain.Character
}
