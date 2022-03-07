package test

import (
	"github.com/Tambarie/movie-api/internal/adapters/api/movie"
	"github.com/Tambarie/movie-api/internal/adapters/server"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPHandler_GetMovies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	mockedRedisService := ports.NewMockRedisService(controller)
	mockedMovieService := ports.NewMockMovieService(controller)
	handler := api.NewHTTPHandler(mockedMovieService, mockedRedisService)

	router := gin.Default()
	server.DefineRouter(handler, router)

	t.Run("Test_for_movie_is_nil", func(t *testing.T) {
		mockedRedisService.EXPECT().GetMovie("movies").Return(nil)
		mockedMovieService.EXPECT().CountComments(gomock.Any()).Return(int64(12), nil).Times(6)
		mockedRedisService.EXPECT().SetMovie("movies", gomock.Any()).Return(nil)
		req, err := http.NewRequest(http.MethodGet, "/api/movies", nil)
		if err != nil {
			t.Fatalf("an error occured: %v", err)
		}
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		if !strings.Contains(response.Body.String(), "Revenge of the Sith\",\"opening_crawl") {
			t.Errorf("Expected Revenge of the Sith to be present")
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected: %v, Got: %v", http.StatusOK, response.Code)
		}
	})
}
