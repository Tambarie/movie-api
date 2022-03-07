package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
		// Getting comments from the database
		movieData, err := h.movieService.GetComments(movieID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		}

		// successful retrieval of comments from DB
		context.JSON(200, gin.H{
			"message": "comments retrieved",
			"data":    movieData,
		})
	}
}
