package ports

import domain "github.com/Tambarie/movie-api/internal/core/domain/movie"

type MovieRepository interface {
	SaveComments(comment *domain.Movie) (*domain.Movie, error)
	GetComments(movieId int) (*[]domain.Comment, error)
	CountComments(movieId int) (int64, error)
}

type RedisRepository interface {
	SetMovie(key string, value *[]domain.Movie) error
	GetMovie(key string) *[]domain.Movie
	SetMovieCharacters(key string, value []domain.Character)
	GetMovieCharacters(key string) []domain.Character
}
