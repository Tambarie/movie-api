package test

import (
	"bytes"
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

func TestAddCommentsToMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := gomock.NewController(t)
	mockedRedisService := ports.NewMockRedisService(controller)
	mockedMovieService := ports.NewMockMovieService(controller)
	handler := api.NewHTTPHandler(mockedMovieService, mockedRedisService)

	router := gin.Default()
	server.DefineRouter(handler, router)

	t.Run("Test_for_adding_comments_to movie", func(t *testing.T) {

		mockedComment := &domain.Comment{}
		mockedComment.IP = "1"
		mockedComment.Content = "Here I am"
		mockedComment.MovieID = 1

		body := []byte(`{"ip":"12","content":"it's emmanuel"}`)

		mov := &[]domain.Movie{
			{
				EpisodeId:    1,
				Title:        "",
				OpeningCrawl: "",
				CommentCount: 0,
				ReleaseDate:  "",
			},
		}
		mockedMovieService.EXPECT().SaveComments(gomock.Any()).Return(mockedComment, nil)
		mockedRedisService.EXPECT().GetMovie("movies").Return(mov)
		mockedRedisService.EXPECT().SetMovie("movies", gomock.Any()).Return(nil)
		req, err := http.NewRequest(http.MethodPost, "/api/movies/1/comments", bytes.NewBuffer(body))

		if err != nil {
			t.Fatalf("an error occured:%v", err)

		}

		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
		responseBody := "successfully added comment to movie"
		responseBody1 := `"movie_id":1,"content":"Here I am","ip":"1"`
		if !strings.Contains(response.Body.String(), responseBody) {
			t.Errorf("Expected body to contain %s", responseBody)
		}
		if !strings.Contains(response.Body.String(), responseBody1) {
			t.Errorf("Expected body to contain %s", responseBody)
		}

		controller.Finish()

	})
}
