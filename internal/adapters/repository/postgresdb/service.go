package postgresdb

import (
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"log"
)

func (r *PostgresRepository) SaveComments(comment *domain.Comment) (*domain.Comment, error) {

	log.Println("comment is about to save")
	err := r.DB.Create(&comment).Error
	log.Println("comment is  saved")
	if err != nil {
		return nil, err
	}
	log.Println("COMMENT HERE", comment)
	return comment, nil
}

func (r *PostgresRepository) GetComments(movieID int) (*[]domain.Comment, error) {
	var comments []domain.Comment
	err := r.DB.Where("movie_id = ?", movieID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return &comments, nil
}

func (r *PostgresRepository) CountComments(movieID int) (int64, error) {
	var counter int64
	err := r.DB.Model(&domain.Comment{}).Where("movie_id = ? ", movieID).Count(&counter).Error
	if err != nil {
		return 0, err
	}
	return counter, nil
}
