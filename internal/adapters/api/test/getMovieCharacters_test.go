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

func TestGetMovieCharacters(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockedRedisService := ports.NewMockRedisService(controller)
	mockedMovieService := ports.NewMockMovieService(controller)
	handler := api.NewHTTPHandler(mockedMovieService, mockedRedisService)

	router := gin.Default()
	server.DefineRouter(handler, router)

	t.Run("get_movies_character_from_redis", func(t *testing.T) {
		character := []domain.Character{
			{
				"Jabba Desilijic Tiure",
				"175",
				"1",
				"358",
				"green-tan, brown",
				"orange",
				"600BBY",
				"hermaphrodite",
			},
		}

		mockedRedisService.EXPECT().GetMovieCharactersInRedis("1").Return(character)
		request, err := http.NewRequest(http.MethodGet, "/api/movies/1/characters", nil)

		if err != nil {
			t.Fatalf("an error occured:%v", err)

		}
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		responseBody := `"skin_color":"green-tan, brown","eye_color":"orange"`

		if response.Code != 201 {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
		if !strings.Contains(response.Body.String(), responseBody) {
			t.Errorf("Expected body to contain %s", responseBody)
		}
	})
}
