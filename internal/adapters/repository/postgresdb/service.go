package postgresdb

import domain "github.com/Tambarie/movie-api/internal/core/domain/movie"

func (r *PostgresRepository) SaveComments(comment *domain.Movie) (*domain.Movie, error) {
	return nil, nil
}

func (r *PostgresRepository) GetComments(movieId int) (*[]domain.Comment, error) {
	return nil, nil
}

func (r *PostgresRepository) CountComments(movieId int) (int64, error) {
	return 0, nil
}
