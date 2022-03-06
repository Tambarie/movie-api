package api

import (
	"fmt"
	"github.com/Tambarie/movie-api/internal/adapters/api/swapi"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/core/shared"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
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

// @Summary      Route Gets all movies
// @Description  List an array of movies containing the name, opening crawl and comment count"
// @Produce  json
// @Success 200 {object} []domain.Movie
// @Router /movies [get]
func (h *HTTPHandler) GetMovies() gin.HandlerFunc {
	return func(context *gin.Context) {

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

			h.AddComment(movies)

			err = h.redisService.SetMovie("movies", movies)
			if err != nil {
				context.JSON(http.StatusInternalServerError, "could not set value to redis")
			}
		}
		context.JSON(http.StatusOK, movies)
	}
}

// @Summary Get characters
// @Description accept sort parameters to sort by one of name, gender or height in ascending or descending order."
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Param sortBy query string  false "Sort by height or name or gender"
// @Param order query string false "descending or ascending order"
// @Param filterBy query string false "can be filtered by male or female options"
// @Success 200 {object} []domain.Character
// @Router /movies/{movieID}/characters/ [get]
func (h *HTTPHandler) GetMoviesCharacters() gin.HandlerFunc {
	return func(context *gin.Context) {
		sorter := context.Query("sort")
		filter := context.Query("filter")
		filter = strings.TrimSpace(filter)
		order := context.Query("order")
		movieIDStr := context.Param("movieID")
		fmt.Println(movieIDStr)
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			//context.JSON(http.StatusInternalServerError, err)
			return
		}

		movieCharacters := h.redisService.GetMovieCharactersInRedis(movieIDStr)

		if movieCharacters == nil {
			movieLinks, err := api.GetAllCharactersByMovieID(movieID)

			if err != nil {
				context.JSON(http.StatusInternalServerError, shared.REQUEST_NOT_FOUND)
				return
			}

			for _, movieLink := range *movieLinks {
				info, err := api.GetCharacterInfo(movieLink)
				if err != nil {
					log.Println(err)
				}

				movieCharacters = append(movieCharacters, *info)
			}
			_ = h.redisService.SetMovieCharactersInRedis(movieIDStr, movieCharacters)
		}

		if filter == "n/a" || filter == "hermaphrodite" || filter == "female" || filter == "male" {
			filteredObject := []domain.Character{}

			for _, movieCharacter := range movieCharacters {
				if movieCharacter.Gender == filter {
					filteredObject = append(filteredObject, movieCharacter)
				}
			}
			movieCharacters = filteredObject
		}

		if order == "ascending" {
			switch sorter {
			case "name":
				sort.Slice(movieCharacters, func(i, j int) bool {
					return movieCharacters[i].Name < movieCharacters[j].Name
				})
			case "gender":
				sort.Slice(movieCharacters, func(i, j int) bool {
					return movieCharacters[i].Gender < movieCharacters[j].Gender
				})

			case "height":
				sort.Slice(movieCharacters, func(i, j int) bool {
					I, _ := strconv.ParseFloat(movieCharacters[i].Height, 64)
					J, _ := strconv.ParseFloat(movieCharacters[j].Height, 64)
					return I < J
				})

			}
		} else {
			switch sorter {
			case "name":
				sort.Slice(movieCharacters, func(i, j int) bool {
					return movieCharacters[i].Name > movieCharacters[j].Name
				})
			case "gender":
				sort.Slice(movieCharacters, func(i, j int) bool {
					return movieCharacters[i].Gender > movieCharacters[j].Gender
				})

			case "height":
				sort.Slice(movieCharacters, func(i, j int) bool {
					I, _ := strconv.ParseFloat(movieCharacters[i].Height, 64)
					J, _ := strconv.ParseFloat(movieCharacters[j].Height, 64)
					return I > J
				})

			}
		}

		var totalHeight float64 = 0
		for _, movieCharacter := range movieCharacters {
			height, err := strconv.ParseFloat(movieCharacter.Height, 64)
			if err != nil {
				log.Println(err)
				continue
			}

			totalHeight += height
		}

		feet := totalHeight / 30.48
		inches := totalHeight / 2.54

		result := fmt.Sprintf("%0.2fcm, %0.2fft , %0.2finches", totalHeight, feet, inches)

		context.JSON(201, gin.H{
			"message":          "retrieved successfully",
			"characters":       len(movieCharacters),
			"totalHeight":      result,
			"listOfCharacters": movieCharacters,
		})

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

// @Summary Adds a new comment to a post
// @Description Adds a new comment to a post with  movieID
// @Accept  json
// @Produce  json
// @Param comment body []domain.Comment true "Comment"
// @Param movie_id path int true "MovieId"
// @Router /movies/:movieID/comments/ [post]
func (h *HTTPHandler) AddCommentToMovies() gin.HandlerFunc {
	return func(context *gin.Context) {
		movieIDStr := context.Param("movieID")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		comment := &domain.Comment{}
		comment.IP = context.ClientIP()
		comment.MovieID = movieID

		err = context.BindJSON(&comment)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		if len(comment.Content) >= 500 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "please, you have exceeded the 500 maximum number of inputs",
			})
			return
		}

		commentData, err := h.movieService.SaveComments(comment)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		if !h.IncrementCommentCountInRedis(movieID) {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "movieID not present",
			})
		}

		context.JSON(200, gin.H{
			"message": "successfully added comment to movie",
			"data":    commentData,
		})
	}
}

// @Summary Endpoint Gets a list of comments
// @Description Endpoint Gets a list of comments for a particular movieID
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Success 200 {object} []domain.Comment
// @Router /movies/:movieID/comments/ [get]
// GetComments method returns all comments for a particular movieID
func (h *HTTPHandler) GetCommentsInMovie() gin.HandlerFunc {
	return func(context *gin.Context) {
		movieIDStr := context.Param("movieID")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		movieData, err := h.movieService.GetComments(movieID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		}

		context.JSON(200, gin.H{
			"message": "comments retrieved",
			"data":    movieData,
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
