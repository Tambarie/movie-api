package api

import (
	"github.com/Tambarie/movie-api/internal/adapters/api/swapi"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
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

func (h *HTTPHandler) GetMovies() gin.HandlerFunc {
	return func(context *gin.Context) {

		var movies = h.redisService.GetMovie("movies")

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

			h.AddCommentOnMovies(movies)

			err = h.redisService.SetMovie("movies", movies)
			if err != nil {
				context.JSON(http.StatusInternalServerError, "could not set value to redis")
			}
		}
		context.JSON(http.StatusOK, movies)
	}
}

func (h *HTTPHandler) AddCommentOnMovies(movies *[]domain.Movie) {
	for idx, movie := range *movies {
		count, _ := h.movieService.CountComments(movie.EpisodeId)
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
