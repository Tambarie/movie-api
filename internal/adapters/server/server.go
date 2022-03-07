package server

import (
	"github.com/Tambarie/movie-api/docs"
	api "github.com/Tambarie/movie-api/internal/adapters/api/movie"
	"github.com/Tambarie/movie-api/internal/core/helper"
	"github.com/Tambarie/movie-api/internal/core/shared"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// DefineRouter Routes
func DefineRouter(handler *api.HTTPHandler, router *gin.Engine) {
	router.Use(helper.LogRequest)
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/api/movies", handler.GetMovies())
	router.GET("/api/movies/:movieID/characters", handler.GetMoviesCharacters())
	router.POST("/api/movies/:movieID/comments", handler.AddCommentToMovies())
	router.GET("/api/movies/:movieID/comments", handler.GetCommentsInMovie())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helper.PrintErrorMessage("404", shared.NoResourceFound))
	})

}
