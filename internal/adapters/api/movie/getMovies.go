package api

import (
	"fmt"
	api "github.com/Tambarie/movie-api/internal/adapters/api/swapi"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

// @Summary      Route Gets all movies
// @Description  List an array of movies containing the name, opening crawl and comment count"
// @Produce  json
// @Success 200 {object} []domain.Movie
// @Router /movies [get]
func (h *HTTPHandler) GetMovies() gin.HandlerFunc {
	return func(context *gin.Context) {

		// Getting movie from the redis DB
		movies := h.redisService.GetMovie("movies")

		if movies == nil {
			data, err := api.GetAllMovies()

			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			result := *data
			sort.Slice(result, func(i, j int) bool {
				return result[i].ReleaseDate > result[j].ReleaseDate
			})
			movies = &result

			// Adding comments to movie
			h.AddComment(movies)

			err = h.redisService.SetMovie("movies", movies)
			if err != nil {
				context.JSON(http.StatusInternalServerError, "could not set value to redis")
				return
			}
		}
		context.JSON(http.StatusOK, movies)
	}
}

//  @Summary AddCounntToMovies
// @Description  a method that add comments to movies in the redis cache
func (h *HTTPHandler) AddComment(movies *[]domain.Movie) {
	for idx, movie := range *movies {
		count, _ := h.movieService.CountComments(movie.EpisodeId)
		fmt.Println(count)
		temp := domain.Movie{
			EpisodeId:    movie.EpisodeId,
			Title:        movie.Title,
			CommentCount: count,
			OpeningCrawl: movie.OpeningCrawl,
			ReleaseDate:  movie.ReleaseDate,
		}
		(*movies)[idx] = temp
	}
}
