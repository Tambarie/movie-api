package api

import (
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Adds a new comment to a post
// @Description Adds a new comment to a post with  movieID
// @Accept  json
// @Produce  json
// @Param comment body []domain.Comment true "Comment"
// @Param movie_id path int true "MovieId"
// @Router /movies/:movieID/comments/ [post]
func (h *HTTPHandler) AddCommentToMovies() gin.HandlerFunc {
	return func(context *gin.Context) {
		//getting the parameter from the url path
		movieIDStr := context.Param("movieID")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		// initializing the comment struct
		comment := &domain.Comment{}
		comment.IP = context.ClientIP()
		comment.MovieID = movieID

		// binding the initialized fields to json
		err = context.BindJSON(&comment)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		// check if the length of the comment is more than 500
		if len(comment.Content) >= 500 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "please, you have exceeded the 500 maximum number of inputs",
			})
			return
		}

		//saving the comments to the DB
		commentData, err := h.movieService.SaveComments(comment)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		// Increment comment in redisDB
		if !h.IncrementCommentCountInRedis(movieID) {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "movieID not present",
			})
		}

		//Successful adding  of comment
		context.JSON(200, gin.H{
			"message": "successfully added comment to movie",
			"data":    commentData,
		})
	}
}

//  @Summary IncrementCommentCountInRedis
// @Description  a method that increments  comment count to movies in the redisDB
func (h *HTTPHandler) IncrementCommentCountInRedis(movieID int) bool {
	var moviesInRedis = h.redisService.GetMovie("movies")

	if moviesInRedis != nil {
		for i, movie := range *moviesInRedis {
			if movie.EpisodeId == movieID {
				incr := domain.Movie{
					EpisodeId:    movie.EpisodeId,
					Title:        movie.Title,
					OpeningCrawl: movie.OpeningCrawl,
					CommentCount: movie.CommentCount + 1,
					ReleaseDate:  movie.ReleaseDate,
				}
				(*moviesInRedis)[i] = incr
				err := h.redisService.SetMovie("movies", moviesInRedis)
				if err != nil {
					return false
				}
				return true
			}
		}
	}
	return false
}
