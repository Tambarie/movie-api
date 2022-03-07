package api

import (
	"fmt"
	api "github.com/Tambarie/movie-api/internal/adapters/api/swapi"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/core/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// @Summary Get characters
// @Description accept sort parameters to sort by one of name, gender or height in ascending or descending order."
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Param sortBy query string  false "Sort by height or name or gender"
// @Param order query string false "descending or ascending order"
// @Param filterBy query string false "can be filtered by male or female options"
// @Success 200 {object} []domain.Character
// @Router /movies/:movieID/characters/ [get]
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

		// Getting movie characters from redis DB
		movieCharacters := h.redisService.GetMovieCharactersInRedis(movieIDStr)

		if movieCharacters == nil {
			//
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

			// Setting movie characters from redis DB
			_ = h.redisService.SetMovieCharactersInRedis(movieIDStr, movieCharacters)
		}

		// Filtering movie characters
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

		// Calculating the total height of movie characters
		var totalHeight float64 = 0
		for _, movieCharacter := range movieCharacters {
			height, err := strconv.ParseFloat(movieCharacter.Height, 64)
			if err != nil {
				log.Println(err)
				continue
			}

			totalHeight += height
		}

		// Converting from cm to feet
		feet := totalHeight / 30.48

		// Converting from cm to inches
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
