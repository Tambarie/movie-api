package test

import (
	api "github.com/Tambarie/movie-api/internal/adapters/api/movie"
	"github.com/Tambarie/movie-api/internal/adapters/server"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockedRedisService := ports.NewMockRedisService(controller)
	mockedMovieService := ports.NewMockMovieService(controller)
	handler := api.NewHTTPHandler(mockedMovieService, mockedRedisService)

	router := gin.Default()
	server.DefineRouter(handler, router)

	t.Run("test for adding comments to movie", func(t *testing.T) {

		comment := &[]domain.Comment{
			{Model: domain.Model{},
				MovieID: 1,
				Content: "This is the comment",
				IP:      "1",
			},
		}
		mockedMovieService.EXPECT().GetComments(1).Return(comment, nil)
		request, err := http.NewRequest(http.MethodGet, "/api/movies/1/comments", nil)
		if err != nil {
			t.Fatalf("an error occured:%v", err)

		}
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		responseBody := `"content":"This is the comment","ip":"1"`

		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

		if !strings.Contains(response.Body.String(), responseBody) {
			t.Errorf("Expected body to contain %s", responseBody)
		}
	})

}
