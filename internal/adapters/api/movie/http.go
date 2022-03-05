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

			h.AddCommentOnMovies(movies)

			err = h.redisService.SetMovie("movies", movies)
			if err != nil {
				context.JSON(http.StatusInternalServerError, "could not set value to redis")
			}
		}
		context.JSON(http.StatusOK, movies)
	}
}

func (h *HTTPHandler) GetMoviesCharacters() gin.HandlerFunc {
	return func(context *gin.Context) {
		sorter := context.Query("sort")
		filter := context.Query("filter")
		filter = strings.TrimSpace(filter)
		order := context.Query("order")
		movieIDStr := context.Param("movieID")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			context.JSON(http.StatusInternalServerError, shared.CONVERSION_ERROR)
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

		if filter == "female" || filter == "male" {
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
			"message":     "retrieved successfully",
			"characters":  len(movieCharacters),
			"totalHeight": result,
		})
	}
}

func (h *HTTPHandler) AddCommentOnMovies(movies *[]domain.Movie) {
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

func (h *HTTPHandler) AddCommentToMovies() gin.HandlerFunc {
	return func(context *gin.Context) {
		movieIDStr := context.Param("movieID")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			context.JSON(http.StatusInternalServerError, shared.CONVERSION_ERROR)
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

		//for i := 0; i < len(*movieData)/2; i++ {
		//	(*movieData)[i], (*movieData)[len(*movieData)-1-i] = movieData[len(*movieData)-1-i], (*movieData)[i]
		//

		context.JSON(200, gin.H{
			"message": "comments retrieved",
			"data":    movieData,
		})
	}
}

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
